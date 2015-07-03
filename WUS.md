# WUS
Wid Url Shortener

## Config

	base_url: 'http://wus.wid.la'
	db_url: '127.0.0.1'
	db_user: 'root'
	db_pass: 'rootpass'
	

## Models

#### Users

	id: 1
	username: toto
	password: bla
	admin: true

#### Providers

	id: 1
	provider: 'wus'
	provider_url: 'http://wus.wid.la'

#### Urls

    id: 1
    provider_key: 1
    user_key: 1
    short_handle: '1afdfa'
    long_url: 'https://docs.angularjs.org/guide/directive'
    notes: 'some nice url'
    enabled: true


## API

/xxx 

	http://wus.wid.la/1afdfa -> redirection

/api

	/user
	/url
	/provider

## Plugins

* Default provider WUS
* goo.gl
* ...