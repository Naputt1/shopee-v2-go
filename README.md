# shopee-v2-go

> Auto-generated, type-safe Go SDK for the Shopee Open API v2

**28 service modules**, **445 API methods**, **840 error constants**, and **428 tests** — all auto-generated from [Shopee's official API documentation](https://open.shopee.com/documents/) using [doclient](https://github.com/Naputt1/doclient).

- HMAC-SHA256 request signing
- Automatic access token refresh on expiry
- Rate-limit retry with `Retry-After`
- OAuth 2.0 auth URL generation
- File upload support
- Go 1.22 generics for extensible metadata

## Install

```bash
go get github.com/naputt1/shopee-v2-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/naputt1/shopee-v2-go"
)

func main() {
    app := goshopee.App{
        PartnerID:  os.Getenv("SHOPEE_PARTNER_ID"),
        PartnerKey: os.Getenv("SHOPEE_PARTNER_KEY"),
        APIURL:     "https://partner.shopeemobile.com",
    }

    client := goshopee.NewDefaultClient(app)

    result, err := client.Product.GetItemList(
        context.Background(),
        123456,                               // shop_id
        goshopee.ProductGetItemListRequest{    // GET query params
            PaginationOffset: goshopee.Ptr(0),
            PaginationLimit:  goshopee.Ptr(10),
        },
        "your_access_token",
    )
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range result.Response.ItemList {
        fmt.Printf("  %s (ID: %d)\n", item.ItemName, item.ItemId)
    }
}
```

## Client

The client is generic over a metadata type `T`.

### Default client

```go
client := goshopee.NewDefaultClient(app)
// client is *goshopee.Client[any]
```

### With custom metadata

```go
type Meta struct {
    ShopID   uint64
    Merchant string
}

client := goshopee.NewClient[Meta](app, goshopee.WithMeta(Meta{
    ShopID:   123456,
    Merchant: "my-store",
}))
```

The `Meta` value is passed to your `OnTokenRefresh` callback, so you can persist the new refresh token alongside your business data:

```go
goshopee.WithOnTokenRefresh(func(res *goshopee.RefreshAccessTokenResponse, meta Meta) {
    db.SaveRefreshToken(meta.ShopID, res.RefreshToken)
}),
```

### Options

| Option | Description |
|--------|-------------|
| `WithHTTPClient` | Custom `*http.Client` (default: 10s timeout) |
| `WithRetry` | Number of retry attempts on failure |
| `WithLogger` | Custom `LeveledLoggerInterface` |
| `WithProxy` | HTTP proxy URL |
| `WithRefreshToken` | Refresh token for auto-renewal |
| `WithOnTokenRefresh` | Callback after successful token refresh |
| `WithMeta` | Custom metadata (generic) |

## Usage

### GET request

```go
result, err := client.Product.GetItemList(ctx, sid, goshopee.ProductGetItemListRequest{
    PaginationOffset: goshopee.Ptr(0),
    PaginationLimit:  goshopee.Ptr(100),
    ItemStatus:       (*goshopee.ItemStatus)(goshopee.Ptr("NORMAL")),
}, tok)
```

### POST request

```go
req := goshopee.AddItemRequest{
    CategoryId:     100001,
    ItemName:       "My Product",
    Description:    "Product description",
    OriginalPrice:  29.99,
    Weight:         0.5,
    Image:          &goshopee.RequestImage{ImageUrl: []string{"https://..."}},
    LogisticInfo:   []goshopee.Logistic{{LogisticId: goshopee.Ptr(int64(20001))}},
    ItemStatus:     (*goshopee.ItemStatus)(goshopee.Ptr("NORMAL")),
}

result, err := client.Product.AddItem(ctx, sid, req, tok)
```

### Media upload

```go
// From file
req, err := client.NewUploadRequest(ctx, "media_space/upload_image", "image", "photo.jpg", sid, tok)

// From io.Reader
req, err := client.NewUploadFromReaderRequest(ctx, "media_space/upload_image", "image", "photo.jpg", reader, sid, tok)
```

### Pointer helpers

The SDK provides a `Ptr` helper for optional fields:

```go
goshopee.Ptr("NORMAL")     // *string
goshopee.Ptr(int64(100))   // *int64
goshopee.Ptr(29.99)        // *float64
```

## Authentication

### OAuth flow

```go
// 1. Generate auth URL
authURL := client.Auth.GetAuthURL()

// 2. User authorizes — Shopee redirects with `code` and `shop_id`

// 3. Exchange code for tokens
token, _ := client.Auth.GetAccessToken(ctx, sid, 0, "code_from_redirect")
// token.AccessToken, token.RefreshToken, token.ExpireIn

// 4. Refresh on expiry (built-in if WithRefreshToken is set)
newToken, _ := client.Auth.RefreshAccessToken(ctx, sid, 0, "refresh_token")
```

### Auto refresh

When `WithRefreshToken` is set, the client automatically detects `error_invalid_access_token` and `error_access_token_expired` errors, refreshes the token, and retries the request — no manual intervention needed.

```go
client := goshopee.NewDefaultClient(app,
    goshopee.WithRefreshToken("initial_refresh_token"),
    goshopee.WithOnTokenRefreshDefault(func(res *goshopee.RefreshAccessTokenResponse, meta any) {
        // Persist the new tokens
    }),
)
```

## Error Handling

### ResponseError

All API calls return `*ResponseError` on failure:

```go
type ResponseError struct {
    Status     int
    Message    string
    Errors     string
    ShopeeError string
    RequestID  string
}
```

### Rate limits

A 429 response returns `*RateLimitError` which embeds `ResponseError` and includes the `RetryAfter` duration.

### Shopee error matching

```go
if goshopee.IsShopeeError(err, goshopee.ErrErrorItemNotFound) {
    // handle "Product not found"
}
```

The SDK includes **840 named error constants** covering all Shopee error codes:

```go
Err10002                                    // error_shop_not_exists
ErrErrorItemNotFound                        // Product not found
ErrErrorInvalidCategory                     // Invalid category ID
ErrErrorParamMissing                        // Missing required parameter
ErrAddOnAddOnDealExpired                    // Expired add-on deal
ErrLogisticsLogisticIsNotAvailable          // Logistic not available
// ... 835 more
```

Use `IsShopeeError(err, code)` to check for a specific Shopee error regardless of the HTTP status code.

## Enum Types

The SDK provides typed string constants for common Shopee enum values:

```go
type ItemStatus string
const (
    ItemStatusNORMAL  ItemStatus = "NORMAL"
    ItemStatusBANNED  ItemStatus = "BANNED"
    ItemStatusUNLIST  ItemStatus = "UNLIST"
)

type OrderStatus string
const (
    OrderStatusUnpaid       OrderStatus = "UNPAID"
    OrderStatusReadyToShip  OrderStatus = "READY_TO_SHIP"
    OrderStatusProcessed    OrderStatus = "PROCESSED"
    OrderStatusShipped      OrderStatus = "SHIPPED"
    OrderStatusCompleted    OrderStatus = "COMPLETED"
    OrderStatusInCancel     OrderStatus = "IN_CANCEL"
    OrderStatusCancelled    OrderStatus = "CANCELLED"
    OrderStatusInvoicePending OrderStatus = "INVOICE_PENDING"
)
```

Other enums: `BookingStatus`, `CampaignStatus`, `DescriptionType`, `DescriptionElementFieldType`, `InvoiceOption`, `LogisticsStatus`, `OperationType`, `PromotionStatus`, `ReturnStatus`, `TaxType`, `WarrantyTime`.

## BoolString

Shopee's API sometimes returns boolean values as strings (`"TRUE"` / `"FALSE"`) and sometimes as JSON booleans. The `BoolString` type handles both:

```go
type BoolString bool

// Unmarshals both "TRUE" and true
var bs goshopee.BoolString
json.Unmarshal([]byte(`"TRUE"`), &bs)  // true
json.Unmarshal([]byte(`true`), &bs)    // true
```

## Services

| Service | Methods | Files |
|---------|---------|-------|
| AccountHealth | 8 | `account_health.gen.go` |
| AddOnDeal | 11 | `add_on_deal.gen.go` |
| Ads | 10 | `ads.gen.go` |
| AMS | 10 | `ams.gen.go` |
| BundleDeal | 1 | `bundle_deal.gen.go` |
| Discount | 14 | `discount.gen.go` |
| FBS | 10 | `fbs.gen.go` |
| FirstMile | 3 | `first_mile.gen.go` |
| FollowPrize | 5 | `follow_prize.gen.go` |
| GlobalProduct | 24 | `global_product.gen.go` |
| Livestream | 8 | `livestream.gen.go` |
| Logistics | 15 | `logistics.gen.go` |
| Media | 2 | `media.gen.go` |
| MediaSpace | 12 | `media_space.gen.go` |
| Merchant | 1 | `merchant.gen.go` |
| Order | 16 | `order.gen.go` |
| Payment | 2 | `payment.gen.go` |
| Product | 51 | `product.gen.go` |
| Public | 2 | `public.gen.go` |
| Push | 7 | `push.gen.go` |
| Returns | 11 | `returns.gen.go` |
| SBS | 3 | `sbs.gen.go` |
| Shop | 16 | `shop.gen.go` |
| ShopCategory | 8 | `shop_category.gen.go` |
| ShopFlashSale | 6 | `shop_flash_sale.gen.go` |
| TopPicks | 6 | `top_picks.gen.go` |
| Video | 4 | `video.gen.go` |
| Voucher | 6 | `voucher.gen.go` |
| Auth | 4 | `auth.go` (hand-written) |

## BaseResponse

Every API response embeds `BaseResponse`:

```go
type BaseResponse struct {
    Error     string `json:"error"`
    Message   string `json:"message"`
    RequestID string `json:"request_id"`
    Warning   string `json:"warning"`
}
```

Check `result.Error` for Shopee-level errors (empty string means success):

```go
result, err := client.Product.AddItem(ctx, sid, req, tok)
if err != nil {
    // HTTP-level error (network, 4xx, 5xx)
}
if result.Error != "" {
    // Shopee API-level error
    fmt.Printf("Shopee error: %s - %s\n", result.Error, result.Message)
}
```

## Regenerating

The SDK is generated using [doclient](https://github.com/Naputt1/doclient). To regenerate from the latest Shopee API docs:

```bash
pnpm install
pnpm run generate
```

See `doclient.config.ts` for the generation configuration.

## Testing

```bash
go test ./...
```

All tests use [httpmock](https://github.com/jarcoal/httpmock) with real API response fixtures (406 JSON files in `fixtures/`). No network access required.

## License

MIT
