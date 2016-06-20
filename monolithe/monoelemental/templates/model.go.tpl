{{ header }}
{% set glob = {'identifier': '', 'prefix': ''} -%}
package {{ package_name }}

{% if package_name != 'elemental' -%}
import "github.com/aporeto-inc/elemental"
{% set _ = glob.update({'prefix': 'elemental.'}) -%}
{% endif %}

{% for imp in imports -%}
import "{{imp}}"
{% endfor -%}

const (
{% for attribute in specification.attributes -%}
    {{ specification.entity_name }}AttributeName{{attribute.local_name[0:1].upper() + attribute.local_name[1:]}} {{ glob.prefix }}AttributeSpecificationNameKey = "{{ specification.rest_name }}/{{ attribute.local_name }}"
{% endfor -%}
)

{% for attr, constant in constants.iteritems() -%}
// {{ constant.type }} represents the possible values for attribute "{{attr}}".
type {{ constant.type }} string
const (
{% for value in constant['values'] -%}
    {{ value.name }} {{ constant.type }} = "{{ value.value }}"
{% endfor -%}
)
{% endfor -%}

// {{specification.entity_name}}Identity represents the Identity of the object
var {{specification.entity_name}}Identity = {{ glob.prefix }}Identity {
    Name:     "{{specification.rest_name}}",
    Category: "{{specification.resource_name}}",
}

{% if not specification.is_root -%}
// {{specification.entity_name_plural}}List represents a list of {{specification.entity_name_plural}}
type {{specification.entity_name_plural}}List []*{{specification.entity_name}}
{%- endif %}

// {{specification.entity_name}} represents the model of a {{specification.rest_name}}
type {{specification.entity_name}} struct {
    {% for attribute in specification.attributes -%}
    {% set field_name = attribute.local_name[0:1].upper() + attribute.local_name[1:] -%}
    {% set json_tags = 'json:"%s,omitempty"' % attribute.local_name if attribute.exposed else 'json:"-"' -%}
    {% set primary_key = ',primarykey' if attribute.primary_key else '' -%}
    {% set cql_tags = 'cql:"%s%s,omitempty"' % (attribute.local_name.lower(), primary_key) if attribute.stored else 'cql:"-"' -%}
    {% set type = attribute.local_type.split(';')[0] -%}
    {% if attribute.name in constants -%}
    {% set type = constants[attribute.name]['type'] -%}
    {%- endif -%}
    {{ field_name }} {{ type }} `{{json_tags}} {{cql_tags}}`
    {% if attribute.identifier -%}
    {% set _ = glob.update({'identifier': field_name}) -%}
    {% endif -%}
    {% endfor -%}
    {%- if specification.is_root %}
    Token string `json:"APIKey,omitempty"`
    Organization string `json:"enterprise,omitempty"`
    {%- endif %}
}

// New{{specification.entity_name}} returns a new *{{specification.entity_name}}
func New{{specification.entity_name}}() *{{specification.entity_name}} {

    return &{{specification.entity_name}}{
        {% for attribute in specification.attributes -%}
        {% set field_name = attribute.local_name[0:1].upper() + attribute.local_name[1:] -%}
        {% if attribute.type == 'external' -%}
        {% set constructor = attribute.local_type.split(';')[1] -%}
        {% if constructor -%}
        {{ field_name }}: {{ constructor }},
        {% endif %}
        {% elif attribute.default_value -%}
        {% set enclosing_format = '"%s"' if attribute.type in ['string', 'enum'] else '%s' -%}
        {{field_name}}: {{ enclosing_format % attribute.default_value}},
        {% endif -%}
        {% endfor -%}
    }
}

// Identity returns the Identity of the object.
func (o *{{specification.entity_name}}) Identity() {{ glob.prefix }}Identity {

    return {{specification.entity_name}}Identity
}

// Identifier returns the value of the object's unique identifier.
func (o *{{specification.entity_name}}) Identifier() string {

    return o.{{ glob.identifier }}
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *{{specification.entity_name}}) SetIdentifier(ID string) {

    o.{{ glob.identifier }} = ID
}

// Validate valides the current information stored into the structure.
func (o *{{specification.entity_name}}) Validate() {{ glob.prefix }}Errors {

    errors := {{ glob.prefix }}Errors{}

    {% for attribute in specification.attributes -%}
    {% set field_name = attribute.local_name[0:1].upper() + attribute.local_name[1:] -%}
    {% set attribute_name = attribute.local_name -%}

    {% if attribute.allowed_choices != None -%}
    if err := {{ glob.prefix }}ValidateStringInList("{{ attribute_name }}", string(o.{{ field_name }}), []string{"{{ attribute.allowed_choices|join('", "') }}"}); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.allowed_chars != None -%}
    if err := {{ glob.prefix }}ValidatePattern("{{ attribute_name }}", o.{{ field_name }}, "{{ attribute.allowed_chars }}"); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.max_length != None -%}
    if err := {{ glob.prefix }}ValidateMaximumLength("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.max_length }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.min_length != None -%}
    if err := {{ glob.prefix }}ValidateMinimumLength("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.min_length }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.max_value != None -%}
    {% if attribute.type == "float" -%}
    if err := {{ glob.prefix }}ValidateMaximumFloat("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.max_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% else -%}
    if err := {{ glob.prefix }}ValidateMaximumInt("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.max_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}
    {% endif -%}

    {% if attribute.min_value != None -%}
    {% if attribute.type == "float" -%}
    if err := {{ glob.prefix }}ValidateMinimumFloat("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.min_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% else -%}
    if err := {{ glob.prefix }}ValidateMinimumInt("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.min_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}
    {% endif -%}

    {% if attribute.required -%}
    {% if attribute.type == "string" -%}
    if err := {{ glob.prefix }}ValidateRequiredString("{{ attribute_name }}", o.{{ field_name }}); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}
    {% endif -%}

    {% endfor -%}
    return errors
}

{% if specification.is_root -%}
// APIKey returns a the API Key
func (o *{{specification.entity_name}}) APIKey() string {

    return o.Token
}

// SetAPIKey sets a the API Key
func (o *{{specification.entity_name}}) SetAPIKey(key string) {

    o.Token = key
}

{% endif -%}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o {{specification.entity_name}}) SpecificationForAttribute(name {{ glob.prefix }}AttributeSpecificationNameKey) {{ glob.prefix }}AttributeSpecification {

  return {{ specification.entity_name }}AttributesMap[name]
}

var {{ specification.entity_name }}AttributesMap = map[{{ glob.prefix }}AttributeSpecificationNameKey]{{ glob.prefix }}AttributeSpecification{
  {% for attribute in specification.attributes -%}
    {{ specification.entity_name }}AttributeName{{attribute.local_name[0:1].upper() + attribute.local_name[1:]}}: {{ glob.prefix }}AttributeSpecification{
      {% if attribute.allowed_chars -%}
      AllowedChars: "{{ attribute.allowed_chars}}",
      {% endif -%}
      {% if attribute.allowed_choices -%}
      AllowedChoices: []string{"{{ attribute.allowed_choices|join('", "') }}"},
      {% else -%}
      AllowedChoices: []string{},
      {% endif -%}
      {% if attribute.autogenerated -%}
      Autogenerated: true,
      {% endif -%}
      {% if attribute.availability -%}
      Availability: "{{ attribute.availability }}",
      {% endif -%}
      {% if attribute.channel -%}
      Channel: "{{ attribute.channel }}",
      {% endif -%}
      {% if attribute.creation_only -%}
      CreationOnly: true,
      {% endif -%}
      {% if attribute.default_order -%}
      DefaultOrder: true,
      {% endif -%}
      {% if attribute.deprecated -%}
      Deprecated: true,
      {% endif -%}
      {% if attribute.exposed -%}
      Exposed: true,
      {% endif -%}
      {% if attribute.filterable -%}
      Filterable: true,
      {% endif -%}
      {% if attribute.foreign_key -%}
      ForeignKey: true,
      {% endif -%}
      {% if attribute.format -%}
      Format: "{{ attribute.format }}",
      {% endif -%}
      {% if attribute.identifier -%}
      Identifier: true,
      {% endif -%}
      {% if attribute.index -%}
      Index: true,
      {% endif -%}
      {% if attribute.max_length -%}
      MaxLength: {{ attribute.max_length }},
      {% endif -%}
      {% if attribute.max_value -%}
      MaxValue: {{ attribute.max_value }},
      {% endif -%}
      {% if attribute.min_length -%}
      MinLength: {{ attribute.min_length }},
      {% endif -%}
      {% if attribute.min_value -%}
      MinValue: {{ attribute.min_value }},
      {% endif -%}
      {% if attribute.local_name -%}
      Name: "{{ attribute.local_name }}",
      {% endif -%}
      {% if attribute.orderable -%}
      Orderable: true,
      {% endif -%}
      {% if attribute.primary_key -%}
      PrimaryKey: true,
      {% endif -%}
      {% if attribute.read_only -%}
      ReadOnly: true,
      {% endif -%}
      {% if attribute.required -%}
      Required: true,
      {% endif -%}
      {% if attribute.stored -%}
      Stored: true,
      {% endif -%}
      {% if attribute.subtype -%}
      SubType: "{{ attribute.subtype }}",
      {% endif -%}
      {% if attribute.transient -%}
      Transient: true,
      {% endif -%}
      {% if attribute.type -%}
      Type: "{{ attribute.type }}",
      {% endif -%}
      {% if attribute.unique -%}
      Unique: true,
      {% endif -%}
      {% if attribute.unique_scope -%}
      UniqueScope: "{{ attribute.unique_scope }}",
      {% endif -%}
    },
  {% endfor -%}
}
