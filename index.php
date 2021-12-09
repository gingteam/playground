<?php

declare(strict_types=1);

interface Go
{
    public function hello(): void;
    public function run(): void;
}

/** @var Go $golang_ffi */
$golang_ffi = FFI::cdef(
    <<<CODE
void hello();
void run();
CODE,
    __DIR__.'/main');

$golang_ffi->hello();
$golang_ffi->run();
