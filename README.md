# alkemy-g6

### Create Product

```sh
curl -X POST localhost:8080/api/v1/products \
-H "Content-Type=application/json" \
-d '{"product_code": "P0026", "description": "Product 1", "height": 10.0, "length": 15.0, "width": 5.0, "weight": 1.0, "expiration_rate": 0.1, "freezing_rate": 0.3, "recommended_freezing_temp": -18.0, "product_type_id": 1, "seller_id": 101}'
```