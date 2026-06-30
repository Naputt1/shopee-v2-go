import type {
  Config,
  IR,
  IREndpoint,
  IRModule,
  IRError,
  IRParam,
  SourceAdapter,
} from '@doclient/core';
import { getModuleDisplayName, toPascalCase, createCachedFetcher } from '@doclient/core';

const BASE = 'https://open.shopee.com/opservice/api/v1/doc';

interface ShopeeModule {
  module_id: number;
  module_name: string;
  type: number;
  items: ShopeeItem[];
}

interface ShopeeItem {
  id: number;
  name: string;
}

interface ShopeeAPIInfo {
  api_name: string;
  api_type: string;
  define: string;
  params: string;
  path: string;
  request_sample: string;
  response_sample: string;
  module_id: number;
  error_list: ShopeeErrorDoc[];
}

interface ShopeeErrorDoc {
  name: string;
  description: string;
}

interface ShopeeParams {
  request_params: ShopeeParam[];
  response_params: ShopeeParam[];
}

interface ShopeeParam {
  name: string;
  type: string;
  description: string;
  required: string;
  children: ShopeeParam[];
}

interface ShopeeSample {
  type: string;
  value: string;
}

function normalizeType(shopeeType: string): string {
  switch (shopeeType) {
    case 'int':
    case 'int32':
    case 'int64':
    case 'timestamp':
      return 'integer';
    case 'float':
    case 'double':
      return 'number';
    case 'boolean':
      return 'boolean';
    case 'string':
      return 'string';
    case 'object':
      return 'object';
    case 'object[]':
      return 'object[]';
    case 'int[]':
    case 'int64[]':
      return 'integer[]';
    case 'string[]':
      return 'string[]';
    default:
      return shopeeType;
  }
}

function mapParam(
  p: ShopeeParam,
  isRequest: boolean,
  typeOverrides?: Record<string, string>,
): IRParam {
  const irType = typeOverrides?.[p.name] ?? normalizeType(p.type);
  return {
    name: p.name,
    type: irType,
    shopeeType: p.type,
    description: p.description ?? '',
    required: isParamRequired(p.required),
    children: (p.children ?? []).map((c) => mapParam(c, isRequest, typeOverrides)),
  };
}

function isParamRequired(req: string): boolean {
  if (!req) return true;
  const lower = req.toLowerCase();
  return lower === 'yes' || lower === 'true';
}

function isGetMethod(requestSample: string): boolean {
  try {
    const samples: ShopeeSample[] = JSON.parse(requestSample);
    return samples.some((s) => s.type === 'cURL' && s.value.includes('GET'));
  } catch {
    return false;
  }
}

function getFixtureContent(responseSample: string): string | undefined {
  try {
    const samples: ShopeeSample[] = JSON.parse(responseSample);
    const json = samples.find((s) => s.type === 'JSON');
    return json?.value;
  } catch {
    return undefined;
  }
}

function cleanFixtureName(name: string): string {
  let cleaned = '';
  for (const ch of name) {
    if (/[a-zA-Z0-9._]/.test(ch)) cleaned += ch;
    else cleaned += '_';
  }
  while (cleaned.includes('__')) cleaned = cleaned.replace(/__/g, '_');
  return cleaned;
}

function toGoErrorName(code: string): string {
  const withUnderscore = code.replace(/[-.]/g, '_');
  const parts = withUnderscore.split('_').filter(Boolean);
  const cameled = parts.map((p) => p.charAt(0).toUpperCase() + p.slice(1)).join('');
  return 'Err' + cameled;
}

function isCommonParam(name: string): boolean {
  return ['partner_id', 'shop_id', 'merchant_id', 'access_token', 'timestamp', 'sign'].includes(
    name,
  );
}

function createFetcher(config: Config) {
  if (config.cacheDir) {
    return createCachedFetcher(config.cacheDir).fetchJSON;
  }
  return async <T = unknown>(url: string): Promise<T> => {
    const res = await fetch(url);
    if (!res.ok) throw new Error(`fetch failed: ${url} (${res.status})`);
    return res.json() as Promise<T>;
  };
}

async function fetchModules(fetchJSON: <T>(url: string) => Promise<T>): Promise<ShopeeModule[]> {
  const body = await fetchJSON<{ modules: ShopeeModule[] }>(`${BASE}/module/?version=2`);
  return body.modules ?? [];
}

async function fetchAPIInfo(
  fetchJSON: <T>(url: string) => Promise<T>,
  apiName: string,
): Promise<ShopeeAPIInfo | null> {
  try {
    return await fetchJSON<ShopeeAPIInfo>(
      `${BASE}/api/?version=2&api_name=${encodeURIComponent(apiName)}`,
    );
  } catch {
    return null;
  }
}

function sleep(ms: number): Promise<void> {
  return new Promise((r) => setTimeout(r, ms));
}

export const shopeeSource: SourceAdapter = {
  name: 'shopee',

  async execute(config: Config): Promise<IR> {
    const typeOverrides = config.mappings?.typeOverrides;
    const fetchJSON = createFetcher(config);

    const modules = await fetchModules(fetchJSON);
    const irModules: IRModule[] = [];
    const allErrors = new Map<string, IRError>();
    const fixtures: Array<{ filename: string; content: string }> = [];

    for (const mod of modules) {
      const displayName = getModuleDisplayName(mod.module_name);
      const moduleName = toPascalCase(displayName);

      const irEndpoints: IREndpoint[] = [];

      for (const item of mod.items) {
        const name = item.name.trim();
        if (!name.startsWith('v2.')) continue;

        await sleep(15);

        const info = await fetchAPIInfo(fetchJSON, name);
        if (!info) continue;

        let requestParams: ShopeeParam[] = [];
        let responseParams: ShopeeParam[] = [];
        if (info.params) {
          try {
            const parsed: ShopeeParams = JSON.parse(info.params);
            requestParams = parsed.request_params ?? [];
            responseParams = parsed.response_params ?? [];
          } catch {
            /* ignore */
          }
        }

        for (const e of info.error_list ?? []) {
          const code = e.name?.trim();
          if (code) {
            const key = toGoErrorName(code);
            if (!allErrors.has(key)) {
              allErrors.set(key, { code, description: e.description ?? '' });
            }
          }
        }

        const isGet = isGetMethod(info.request_sample);
        const apiPath = info.path ?? '';
        const relPath = apiPath.replace(/^\/api\/v2\//, '').replace(/^\//, '');

        const parts = name.split('.');
        const baseName = parts.length >= 3 ? parts.slice(2).join('_') : parts[parts.length - 1];
        const methodName = toPascalCase(baseName);
        const fixtureName = cleanFixtureName(name) + '_resp.json';

        const fixtureContent = getFixtureContent(info.response_sample);
        if (fixtureContent) {
          fixtures.push({ filename: fixtureName, content: fixtureContent });
        }

        const irEp: IREndpoint = {
          name: methodName,
          method: isGet ? 'GET' : 'POST',
          path: relPath,
          fullPath: '/api/v2/' + relPath,
          description: info.define ?? '',
          docUrl: `https://open.shopee.com/documents/v2/${name}?module=${mod.module_id}&type=${mod.type}`,
          apiType: (info.api_type as 'Shop' | 'Merchant' | 'Public') || 'Shop',
          isUpload: info.path?.includes('upload') ?? false,
          fullApiName: name,
          requestParams: requestParams
            .filter((p) => !isCommonParam(p.name))
            .map((p) => mapParam(p, true, typeOverrides))
            .sort((a, b) => a.name.localeCompare(b.name)),
          responseParams: responseParams
            .map((p) => mapParam(p, false, typeOverrides))
            .sort((a, b) => a.name.localeCompare(b.name)),
          errors: (info.error_list ?? [])
            .filter((e) => e.name?.trim())
            .map((e) => ({ code: e.name.trim(), description: e.description ?? '' }))
            .sort((a, b) => a.code.localeCompare(b.code)),
        };

        irEndpoints.push(irEp);
      }

      if (irEndpoints.length > 0) {
        irEndpoints.sort((a, b) => a.name.localeCompare(b.name));
        irModules.push({ name: moduleName, moduleId: mod.module_id, endpoints: irEndpoints });
      }
    }

    irModules.sort((a, b) => a.name.localeCompare(b.name));

    const enumsDef = config.mappings?.enums;
    const constants: Array<{
      typeName: string;
      baseType: string;
      values: Array<{ name: string; value: string }>;
    }> = [];
    if (enumsDef) {
      for (const [typeName, def] of Object.entries(enumsDef)) {
        const values: Array<{ name: string; value: string }> = [];
        for (const [value, name] of Object.entries(def.values)) {
          values.push({ name, value });
        }
        values.sort((a, b) => a.name.localeCompare(b.name));
        constants.push({ typeName, baseType: def.base, values });
      }
      constants.sort((a, b) => a.typeName.localeCompare(b.typeName));
    }

    return {
      name: config.name,
      modules: irModules,
      constants,
      errors: Array.from(allErrors.values()).sort((a, b) => a.code.localeCompare(b.code)),
      fixtures: fixtures.sort((a, b) => a.filename.localeCompare(b.filename)),
    };
  },
};
