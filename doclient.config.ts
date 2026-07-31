import { defineConfig } from '@doclient/cli';
import { createGoRenderer } from '@doclient/renderer-go';
import { shopeeSource } from './source-shopee';
import { shopeeProfile } from './profile.js';

export default defineConfig({
  name: 'go-shopee-v2',
  source: shopeeSource,
  output: createGoRenderer(shopeeProfile, {
    package: 'goshopee',
    module: 'github.com/naputt1/shopee-v2-go',
  }),
  cacheDir: '.cache/shopee',
  outputDir: '.',
  mappings: {
    typeOverrides: {
      is_cnsc: 'bool',
      item_status: 'ItemStatus',
      description_type: 'DescriptionType',
      field_type: 'DescriptionElementFieldType',
      warranty_time: 'WarrantyTime',
      invoice_option: 'InvoiceOption',
      tax_type: 'TaxType',
      operation_type: 'OperationType',
      promotion_status: 'PromotionStatus',
      campaign_status: 'CampaignStatus',
      order_status: 'OrderStatus',
      logistics_status: 'LogisticsStatus',
      return_status: 'ReturnStatus',
      booking_status: 'BookingStatus',
      is_eligible: 'bool',
      is_activated: 'bool',
      has_next_page: 'bool',
    },
    structTypeOverrides: {
      GetItemBaseInfoResponseDataItem: { deboost: 'BoolString' },
      ShopeeStock: { stock: 'int64' },
      Item: { product_location_id: 'StringSlice' },
    },
    enums: {
      ItemStatus: {
        base: 'string',
        values: {
          NORMAL: 'ItemStatusNormal',
          BANNED: 'ItemStatusBanned',
          UNLIST: 'ItemStatusUnlist',
        },
      },
      PromotionStatus: {
        base: 'string',
        values: {
          upcoming: 'PromotionStatusUpcoming',
          ongoing: 'PromotionStatusOngoing',
          expired: 'PromotionStatusExpired',
          all: 'PromotionStatusAll',
        },
      },
      CampaignStatus: {
        base: 'string',
        values: {
          upcoming: 'CampaignStatusUpcoming',
          ongoing: 'CampaignStatusOngoing',
          expired: 'CampaignStatusExpired',
          paused: 'CampaignStatusPaused',
          scheduled: 'CampaignStatusScheduled',
          ended: 'CampaignStatusEnded',
          deleted: 'CampaignStatusDeleted',
          closed: 'CampaignStatusClosed',
        },
      },
      OrderStatus: {
        base: 'string',
        values: {
          UNPAID: 'OrderStatusUnpaid',
          READY_TO_SHIP: 'OrderStatusReadyToShip',
          PROCESSED: 'OrderStatusProcessed',
          SHIPPED: 'OrderStatusShipped',
          COMPLETED: 'OrderStatusCompleted',
          IN_CANCEL: 'OrderStatusInCancel',
          CANCELLED: 'OrderStatusCancelled',
          INVOICE_PENDING: 'OrderStatusInvoicePending',
        },
      },
      DescriptionType: {
        base: 'string',
        values: { normal: 'DescriptionTypeNormal', extended: 'DescriptionTypeExtended' },
      },
      DescriptionElementFieldType: {
        base: 'string',
        values: {
          text: 'DescriptionElementFieldTypeText',
          image: 'DescriptionElementFieldTypeImage',
        },
      },
      WarrantyTime: {
        base: 'string',
        values: {
          ONE_YEAR: 'WarrantyTimeOneYear',
          TWO_YEARS: 'WarrantyTimeTwoYears',
          OVER_TWO_YEARS: 'WarrantyTimeOverTwoYears',
        },
      },
      InvoiceOption: {
        base: 'string',
        values: {
          NO_INVOICE: 'InvoiceOptionNoInvoice',
          VAT_MARGIN_SCHEME_INVOICES: 'InvoiceOptionVatMarginScheme',
          VAT_INVOICES: 'InvoiceOptionVat',
          NON_VAT_INVOICES: 'InvoiceOptionNonVat',
        },
      },
      TaxType: {
        base: 'int',
        values: { '0': 'TaxTypeNoTax', '1': 'TaxTypeTaxable', '2': 'TaxTypeTaxFree' },
      },
      OperationType: {
        base: 'string',
        values: { '1': 'OperationTypeRetailer', '2': 'OperationTypeManufactorer' },
      },
      ReturnStatus: {
        base: 'string',
        values: {
          REQUESTED: 'ReturnStatusRequested',
          ACCEPTED: 'ReturnStatusAccepted',
          CANCELLED: 'ReturnStatusCancelled',
          JUDGING: 'ReturnStatusJudging',
          REFUNDING: 'ReturnStatusRefunding',
          CLOSED: 'ReturnStatusClosed',
          PROCESSING: 'ReturnStatusProcessing',
          SELLER_DISPUTE: 'ReturnStatusSellerDispute',
        },
      },
      BookingStatus: {
        base: 'string',
        values: {
          READY_TO_SHIP: 'BookingStatusReadyToShip',
          PROCESSED: 'BookingStatusProcessed',
          SHIPPED: 'BookingStatusShipped',
          CANCELLED: 'BookingStatusCancelled',
          MATCHED: 'BookingStatusMatched',
        },
      },
      LogisticsStatus: {
        base: 'string',
        values: {
          LOGISTICS_NOT_START: 'LogisticsStatusNotStart',
          LOGISTICS_REQUEST_CREATED: 'LogisticsStatusRequestCreated',
          LOGISTICS_PICKUP_DONE: 'LogisticsStatusPickupDone',
          LOGISTICS_PICKUP_FAILED: 'LogisticsStatusPickupFailed',
          LOGISTICS_DELIVERY_DONE: 'LogisticsStatusDeliveryDone',
          LOGISTICS_DELIVERY_FAILED: 'LogisticsStatusDeliveryFailed',
          LOGISTICS_REQUEST_CANCELED: 'LogisticsStatusRequestCanceled',
          LOGISTICS_COD_REJECTED: 'LogisticsStatusCodRejected',
          LOGISTICS_READY: 'LogisticsStatusReady',
          LOGISTICS_INVALID: 'LogisticsStatusInvalid',
          LOGISTICS_LOST: 'LogisticsStatusLost',
        },
      },
    },
    ignoreAPIs: ['v2.public.get_access_token', 'v2.public.refresh_access_token'],
  },
});
