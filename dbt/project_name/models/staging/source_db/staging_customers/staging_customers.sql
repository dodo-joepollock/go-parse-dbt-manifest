SELECT *
FROM {{ source('source_db','customers')}}