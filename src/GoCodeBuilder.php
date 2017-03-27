<?php

namespace Swaggest\GoCodeBuilder;

use Swaggest\CodeBuilder\CodeBuilder;

class GoCodeBuilder extends CodeBuilder
{
    const TRACE_STORAGE_NAME = 'go-code-builder';

    public $goFiles = array();
    public $headComment = <<<GO
// ATTENTION! This file was generated automatically, any manual modifications may be lost

GO;

    public function setBuilderVersion($builderVersion)
    {
        $this->versionComment = <<<GO
// {$builderVersion}

GO;
    }


    public $versionComment;


    public function storeToDisk($srcPath)
    {
        if (!$srcPath) {
            throw new Exception('No src path');
        }

        $srcPath = $this->realSrcPath($srcPath);
        $this->originalFiles = $this->recursiveFileList($srcPath, array('go' => true, 'yaml' => true, 'lock' => true),
            array('vendor', '.git', '*_test.go'));
        parent::storeToDisk($srcPath);
    }


    protected function toCamelCase($string, $lowerFirst = false)
    {
        $result = implode('', array_map('ucfirst', explode('_', $string)));
        if (!$result) {
            return '';
        }
        if ($lowerFirst) {
            $result[0] = strtolower($result[0]);
        }
        return $result;
    }


    public function exportableName($name)
    {
        $goName = preg_replace("/([^a-zA-Z0-9_]+)/", "_", $name);
        $goName = $this->toCamelCase($goName, false);
        if ($goName === 'String') {
            $goName = 'StringProperty';
        }
        if (!$goName) {
            $goName = 'Property' . substr(md5($name), 0, 6);
        } elseif (is_numeric($goName[0])) {
            $goName = 'Property' . $goName;
        }
        return $goName;
    }

    public function privateName($name)
    {
        $goName = preg_replace("/([^a-zA-Z0-9_]+)/", "_", $name);
        $goName = $this->toCamelCase($goName, true);
        if ($goName === 'type') {
            $goName .= substr(md5($name), 0, 6);
        }
        if (!$goName) {
            $goName = 'property' . substr(md5($name), 0, 6);
        } elseif (is_numeric($goName[0])) {
            $goName = 'property' . $goName;
        }
        return $goName;
    }

}