<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome'); // this call view welcome.blade.php
});
