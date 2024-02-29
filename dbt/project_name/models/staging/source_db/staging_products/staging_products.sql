SELECT *
FROM {{ source('source_db','products')}}