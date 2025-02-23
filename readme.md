## Babel Webpack Starter - 2025 Edition
###### This project is an upgraded version of an older Babel and Webpack starter pack, brought up to date with the latest standards and best practises for 2025.

### Dependencies

#### Babel Packages

##### @babel/core: 
###### The core of Babel, used for transforming your JavaScript code.

##### @babel/polyfill: 
###### Provides polyfills necessary for a full ES6+ environment.

##### @babel/preset-env: 
###### A smart preset that allows you to use the latest JavaScript without needing to micromanage which syntax transforms (and optionally, browser polyfills) are needed by your target environments.

##### Various Babel Plugins: 
###### Extend Babel's functionality to support proposals from different stages:

###### @babel/plugin-proposal-function-bind
###### @babel/plugin-proposal-export-default-from
###### @babel/plugin-proposal-logical-assignment-operators
###### @babel/plugin-proposal-optional-chaining
###### @babel/plugin-proposal-pipeline-operator
###### @babel/plugin-proposal-nullish-coalescing-operator
###### @babel/plugin-proposal-do-expressions
###### @babel/plugin-proposal-decorators
###### @babel/plugin-proposal-function-sent
###### @babel/plugin-proposal-export-namespace-from
###### @babel/plugin-proposal-numeric-separator
###### @babel/plugin-proposal-throw-expressions
###### @babel/plugin-syntax-dynamic-import
###### @babel/plugin-syntax-import-meta
###### @babel/plugin-proposal-class-properties
###### @babel/plugin-proposal-json-strings

#### Webpack Packages

##### webpack: 
###### A static module bundler for JavaScript applications.

##### webpack-cli: 
###### Command-line interface for Webpack, providing commands like build and serve.

##### webpack-dev-server: 
###### A development server that provides live reloading, making it easy to develop your applications.

#### Other Packages

##### babel-loader: 
###### Integrates Babel and Webpack, allowing you to use Babel to transpile your JavaScript files during the Webpack build process.

#### Development Configuration
###### Ensure you have the following configuration files:

##### webpack.config.js: 
###### Configures Webpack for bundling and serving your project.

##### babel.config.js: 
###### Configures Babel for transpiling your JavaScript code.

###### This setup allows you to utilise modern JavaScript features, transpile your code for compatibility, and bundle your application efficiently.

#### To build:
##### npm run build

#### To run:
##### npm start

#### Author:
##### Nick Rabone, based off the 2018 Babel-Webpack Starter by Brad Traversy.
