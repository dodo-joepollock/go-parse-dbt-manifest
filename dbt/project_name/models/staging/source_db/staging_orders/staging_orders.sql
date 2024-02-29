SELECT *
FROM {{ source('source_db','orders')}}