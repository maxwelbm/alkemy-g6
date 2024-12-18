# alkemy-g6

## api/v1/products/**

### Get Products

REQ:

```sh
curl localhost:8080/api/v1/products
```

RES:

```sh
{"data":
    [
        {"id":3,"product_code":"P003","description":"Product 3","height":15,"length":10,"width":8,"weight":1.5,"expiration_rate":0.15,"freezing_rate":0.25,"recommended_freezing_temp":-15,"product_type_id":3,"seller_id":103},
        {"id":6,"product_code":"P006","description":"Product 6","height":13,"length":14,"width":5.5,"weight":1.1,"expiration_rate":0.09,"freezing_rate":0.18,"recommended_freezing_temp":-19,"product_type_id":1,"seller_id":106},
        {"id":12,"product_code":"P012","description":"Product 12","height":19,"length":25,"width":10,"weight":2.6,"expiration_rate":0.18,"freezing_rate":0.27,"recommended_freezing_temp":-18,"product_type_id":1,"seller_id":112},
        {"id":16,"product_code":"P016","description":"Product 16","height":25,"length":30,"width":12,"weight":3.5,"expiration_rate":0.13,"freezing_rate":0.26,"recommended_freezing_temp":-22,"product_type_id":3,"seller_id":116},
    ]
}
```

### Get Product by ID

REQ:

```sh
curl localhost:8080/api/v1/products/3
```

RES:

```sh
{"data":
    {"id":3,"product_code":"P003","description":"Product 3","height":15,"length":10,"width":8,"weight":1.5,"expiration_rate":0.15,"freezing_rate":0.25,"recommended_freezing_temp":-15,"product_type_id":3,"seller_id":103}
}
```

### Create Product

REQ:

```sh
curl -X POST localhost:8080/api/v1/products \
-H "Content-Type=application/json" \
-d '{"product_code": "P0026", "description": "Product 1", "height": 10.0, "length": 15.0, "width": 5.0, "weight": 1.0, "expiration_rate": 0.1, "freezing_rate": 0.3, "recommended_freezing_temp": -18.0, "product_type_id": 1, "seller_id": 101}'
```

RES:

```sh
{
    "message":"Created",
    "data":{"ID":26,"ProductCode":"P001","Description":"Product 1","Height":10,"Length":15,"Width":5,"Weight":1,"ExpirationRate":0.1,"FreezingRate":0.3,"RecomFreezTemp":-18,"ProductTypeID":1,"SellerID":0}
}
```
