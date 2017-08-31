<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\JsonSchema\JsonSchema;
use Swaggest\JsonSchema\Schema;

/**
 * @todo properly process $ref, $schema property names
 */
class GoBuilder
{
    /** @var \SplObjectStorage|GeneratedClass[] */
    private $generatedClasses;
    private $untitledIndex = 0;

    public function __construct()
    {
        $this->generatedClasses = new \SplObjectStorage();
    }

    public $buildGetters = false;
    public $buildSetters = false;

    /**
     * @param JsonSchema $schema
     * @param string $path
     * @return AnyType
     */
    public function getType(JsonSchema $schema, $path = '#')
    {
        $typeBuilder = new TypeBuilder($schema, $path, $this);
        return $typeBuilder->build();
    }


    public function getClass(JsonSchema $schema, $path)
    {
        if ($this->generatedClasses->contains($schema)) {
            return $this->generatedClasses[$schema]->class;
        } else {
            return $this->makeClass($schema, $path)->class;
        }
    }

    private function makeClass(Schema $schema, $path)
    {
        if (empty($path)) {
            throw new Exception('Empty path');
        }
        $generatedClass = new GeneratedClass();
        $generatedClass->schema = $schema;

        $class = new PhpClass();
        $class->setName('Untitled' . ++$this->untitledIndex);
        $class->setExtends(Palette::classStructureClass());


        $setupProperties = new PhpFunction('setUpProperties');
        $setupProperties
            ->setVisibility(PhpFlags::VIS_PUBLIC)
            ->setIsStatic(true);
        $setupProperties
            ->addArgument(new PhpNamedVar('properties', Palette::propertiesOrStaticClass()))
            ->addArgument(new PhpNamedVar('ownerSchema', Palette::schemaClass()));

        $body = new PhpCode();

        $class->addMethod($setupProperties);

        $generatedClass->class = $class;
        $generatedClass->path = $path;

        $this->generatedClasses->attach($schema, $generatedClass);

        foreach ($schema->properties as $name => $property) {
            $propertyName = PhpCode::makePhpName($name);

            $schemaBuilder = new SchemaBuilder($property, '$properties->' . $propertyName, $path . '->' . $name, $this);
            $phpProperty = new PhpClassProperty($propertyName, $this->getType($property, $path . '->' . $name));
            if ($property->description) {
                $phpProperty->setDescription($property->description);
            }
            $class->addProperty($phpProperty);
            if ($this->buildGetters) {
                $class->addMethod(new Getter($phpProperty));
            }
            if ($this->buildSetters) {
                $class->addMethod(new Setter($phpProperty));
            }
            $body->addSnippet(
                $schemaBuilder->build()
            );
            if ($propertyName != $name) {
                $body->addSnippet('$ownerSchema->addPropertyMapping(' . var_export($name, 1) . ', self::names()->'
                    . $propertyName . ");\n");
            }

            $class->addMethod(new Setter($phpProperty, true));
        }

        $schemaBuilder = new SchemaBuilder($schema, '$ownerSchema', $path, $this);
        $schemaBuilder->setSkipProperties(true);
        $body->addSnippet($schemaBuilder->build());

        $setupProperties->setBody($body);

        return $generatedClass;
    }

    /**
     * @return GeneratedClass[]
     */
    public function getGeneratedClasses()
    {
        $result = array();
        foreach ($this->generatedClasses as $schema) {
            $result[] = $this->generatedClasses[$schema];
        }
        return $result;
    }

}