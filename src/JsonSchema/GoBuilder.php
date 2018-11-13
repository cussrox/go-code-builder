<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\GoCodeBuilder;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Struct\Tags;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\JsonSchema\JsonSchema;
use Swaggest\JsonSchema\Schema;

class GoBuilder
{
    private $code;

    /** @var Options */
    public $options;

    /** @var GeneratedStruct[] */
    private $generatedStructs;

    /** @var \SplObjectStorage */
    private $generatedStructsBySchema;

    private $untitledIndex = 0;
    /** @var GoCodeBuilder */
    public $codeBuilder;
    public $requiresMarshalMapHelper = false;

    /** @var GoBuilderStructHook */
    public $structPreparedHook;

    /** @var GoBuilderStructHook */
    public $structCreatedHook;

    /** @var MarshalUnion */
    public $marshalUnion;

    /** @var UnmarshalUnion */
    public $unmarshalUnion;

    public function __construct()
    {
        $this->code = new Code();
        $this->options = new Options();
        $this->generatedStructs = [];
        $this->generatedStructsBySchema = new \SplObjectStorage();
        $this->codeBuilder = new GoCodeBuilder();
    }

    public function getCode()
    {
        return $this->code;
    }

    /**
     * @param JsonSchema $schema
     * @param string $path
     * @return AnyType
     * @throws Exception
     */
    public function getType($schema, $path = '#')
    {
        return (new TypeBuilder($schema, $path, $this))->build();
    }

    /**
     * @param $schema
     * @param $path
     * @return StructDef
     * @throws Exception
     */
    public function getClass($schema, $path)
    {
        return $this->getGeneratedStruct($schema, $path)->structDef;
    }

    /**
     * @param $schema
     * @param $path
     * @return mixed|GeneratedStruct
     * @throws Exception
     */
    public function getGeneratedStruct($schema, $path)
    {
        if (isset($this->generatedStructs[$path])) {
            return $this->generatedStructs[$path];
        }

        if ($this->generatedStructsBySchema->contains($schema)) {
            return $this->generatedStructsBySchema[$schema];
        }

        return $this->makeStruct($schema, $path);
    }


    /**
     * @param Schema $schema
     * @param $path
     * @return GeneratedStruct
     * @throws Exception
     */
    private function makeStruct(Schema $schema, $path)
    {
        if (empty($path)) {
            throw new Exception('Empty path');
        }
        $generatedStruct = new GeneratedStruct();
        $this->generatedStructs[$path] = $generatedStruct;
        $this->generatedStructsBySchema->attach($schema, $generatedStruct);
        $generatedStruct->schema = $schema;

        if ($path === '#') {
            $structDef = new StructDef('Untitled' . ++$this->untitledIndex);
        } else {
            $structDef = new StructDef($this->codeBuilder->exportableName($path));
        }

        if ($this->structCreatedHook !== null) {
            $this->structCreatedHook->process($structDef, $path, $schema);
        }

        $comment = $structDef->getName() . ' structure is generated from ' . $path;
        if ($schema->title) {
            $comment .= "\n" . $schema->title;
        }
        if ($schema->description) {
            $comment .= "\n" . $schema->description;
        }
        $structDef->setComment($comment);
        $marshalJson = new MarshalJson($this, $structDef);

        $generatedStruct->structDef = $structDef;
        $generatedStruct->path = $path;
        $generatedStruct->marshalJson = $marshalJson;

        if ($schema->properties !== null) {
            foreach ($schema->properties as $name => $property) {
                $fieldName = $this->codeBuilder->exportableName($name);

                if ($this->options->trimParentFromPropertyNames) {
                    if (strpos($fieldName, $structDef->getName()) === 0) {
                        $fieldName = substr($fieldName, strlen($structDef->getName()));
                    }
                }

                $goPropertyType = $this->getType($property, $path . '->' . $name);
                $goProperty = new StructProperty(
                    $fieldName,
                    $goPropertyType,
                    (new Tags())->setTag('json', $name . ',omitempty')
                );
                $comment = '';
                if ($property->title) {
                    $comment = $property->title . "\n";
                }
                if ($property->description) {
                    $comment .= $property->description;
                }
                if ($comment !== '') {
                    $goProperty->setComment($comment);
                }
                if ($this->options->hideConstProperties) {
                    $enum = [];

                    if ($property->enum !== null) {
                        $enum = $property->enum;
                    }

                    if ($property->const !== null) {
                        $enum = [$property->const];
                    }

                    if (count($enum) === 1) {
                        $marshalJson->constValues[$name] = $enum[0];
                        continue;
                    }
                }

                $structDef->addProperty($goProperty);
                $marshalJson->addNamedProperty($name);
            }
        }

        $structDef->getCode()->addSnippet($marshalJson);

        if ($this->structPreparedHook !== null) {
            $this->structPreparedHook->process($structDef, $path, $schema);
        }

        return $generatedStruct;
    }

    /**
     * @return GeneratedStruct[]
     */
    public function getGeneratedStructs()
    {
        return $this->generatedStructs;
    }

    public function pathToName($path)
    {
        if (0 === strpos($path, '#/definitions/')) {
            return substr($path, strlen('#/definitions/'));
        }

        return $path;
    }

}