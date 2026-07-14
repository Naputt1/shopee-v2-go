import { defineProfile, loadTemplate } from '@doclient/renderer-go';

const clientTpl = loadTemplate('./templates/client.go');
const authTpl = loadTemplate('./templates/auth.go');

export const shopeeProfile = defineProfile({
  name: 'shopee',

  responseDataFieldName: 'response',
  commonFields: ['request_id', 'error', 'message', 'warning'],
  commonRequestFields: [
    'partner_id',
    'shop_id',
    'merchant_id',
    'access_token',
    'timestamp',
    'sign',
  ],

  baseResponseFields: [
    { name: 'Error', type: 'string', jsonTag: 'error', urlTag: '', comment: '' },
    { name: 'Message', type: 'string', jsonTag: 'message', urlTag: '', comment: '' },
    { name: 'RequestID', type: 'string', jsonTag: 'request_id', urlTag: '', comment: '' },
    { name: 'Warning', type: 'string', jsonTag: 'warning,omitempty', urlTag: '', comment: '' },
  ],

  extraMethodParams: ', sid uint64, mid uint64, tok string',
  extraMethodArgs: 'sid, mid, tok',

  renderClientFile: (pkg, services, init) =>
    clientTpl.render({
      PACKAGE_NAME: pkg,
      SERVICES_SECTION: services,
      SERVICES_INIT_SECTION: init,
    }),

  renderAuthFile: (pkg) => authTpl.render({ PACKAGE_NAME: pkg }),

  testSetup: {
    appVars:
      '\tshopID      uint64 = 123456\n\tmerchantID  uint64 = 789012\n\taccessToken        = "test_access_token"',
    appLiteral:
      'App{\n\t\tPartnerID:  123456,\n\t\tPartnerKey: "test_partner_key",\n\t\tRedirectURL: "https://example.com/callback",\n\t\tAPIURL:     "https://open-api.test.com",\n\t}',
  },

  dependencies: ['github.com/google/go-querystring v1.1.0', 'github.com/jarcoal/httpmock v1.3.1'],
});
