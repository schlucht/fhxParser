# Die Datenbank

## Tabelle: units

|  Column |  Type | Comment  |
|---|---|---|
|  id | int  | int |
| type_id |	int	 |  |
| plant_id | int | |
| unit_name |	varchar(255) | NULL	|
| position |	varchar(100) | NULL	|
| time |	int | NULL	|
| author |	varchar(255) |	|
| description |	varchar(255) |	|
| created_at |	datetime | [CURRENT_TIMESTAMP] |	
| updated_at |	datetime | [CURRENT_TIMESTAMP]	|

<hr>

## Tabelle: plants

|  Column |  Type | Comment  |
|---|---|---|
|  id | int  | int |
|  plant_name | varchar(255)  | int |
| created_at |	datetime | [CURRENT_TIMESTAMP] |	
| updated_at |	datetime | [CURRENT_TIMESTAMP]	|
<hr>

## Tabelle: types
|  Column |  Type | Comment  |
|---|---|---|
|  id | int  | int |
|  type_name | varchar(255)  | int |
| created_at |	datetime | [CURRENT_TIMESTAMP] |	
| updated_at |	datetime | [CURRENT_TIMESTAMP]	|

## Tabelle: recipes
|  Column |  Type | Comment  |
|---|---|---|
|  id | int  | int |
|  recipe_name | varchar(255)  |  |
|  author | varchar(255)  |  |
|  description | varchar(255)  |  |
|  unit_id | int  |  |
| created_at |	datetime | [CURRENT_TIMESTAMP] |	
| updated_at |	datetime | [CURRENT_TIMESTAMP]	|

## Tabelle: parameters
|  Column |  Type | Comment  |
|---|---|---|
|  param_id | int  | int |
|  parameter_name | varchar(255)  |  |
|  description | varchar(255)  |  |
|  unit_id | int  |  |
| created_at |	datetime | [CURRENT_TIMESTAMP] |	
| updated_at |	datetime | [CURRENT_TIMESTAMP]	|

## Tabelle: paramvalues
|  Column |  Type | Comment  |
|---|---|---|
|  value_id | int  | int |
|  stringvalue | varchar(255)  |  |
|  value_set | varchar(100)  |  |
|  high | int  |  |
|  low | int  |  |
|  cv | int  |  |
|  unit | varchar(255)  |  |
|  description | varchar(255)  |  |
|  param_id | int  |  |
| created_at |	datetime | [CURRENT_TIMESTAMP] |	
| updated_at |	datetime | [CURRENT_TIMESTAMP]	|