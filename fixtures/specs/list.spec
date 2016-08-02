{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "The description",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "description",
            "orderable": true,
            "read_only": false,
            "required": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "The name",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "getter": true,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "setter": true,
            "orderable": true,
            "read_only": false,
            "required": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "A read only attribute",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "getter": true,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "readOnly",
            "setter": false,
            "orderable": true,
            "read_only": true,
            "required": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "A creation only only attribute",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "getter": true,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "creationOnly",
            "setter": false,
            "orderable": true,
            "read_only": false,
            "required": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "An unexposed attribute",
            "exposed": false,
            "filterable": true,
            "format": "free",
            "getter": true,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "unexposed",
            "setter": false,
            "orderable": true,
            "read_only": false,
            "required": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": true,
            "delete": false,
            "deprecated": false,
            "get": true,
            "relationship": "child",
            "rest_name": "task",
            "update": false
        },
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": false,
            "get": true,
            "relationship": "member",
            "rest_name": "user",
            "update": true
        }
    ],
    "model": {
        "create": false,
        "delete": true,
        "description": "Represent a a list of task to do.",
        "entity_name": "List",
        "extends": [
          "@base"
        ],
        "get": true,
        "package": "todo-list",
        "resource_name": "lists",
        "rest_name": "list",
        "root": false,
        "update": true
    }
}
