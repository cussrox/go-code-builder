<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Schema;

class JsonSchemaGenerateTest extends \PHPUnit_Framework_TestCase
{
    public function testJsonSchemaGenerate()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../../vendor/swaggest/json-schema/spec/json-schema-draft7.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->options->hideConstProperties = true;
        $builder->options->trimParentFromPropertyNames = true;
        $builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder) {
            if ('#' === $path) {
                $structDef->setName('Schema');
            } elseif (0 === strpos($path, '#/definitions/')) {
                $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
                $structDef->setName($name);
            }
        });
        $builder->getType($schema);

        $goFile = new GoFile('entities');
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $expectedGen = file_get_contents(__DIR__ . '/../../../resources/go/draft7/entities.go');

        $this->assertSame($expectedGen, $goFile->render());

    }

}