# demo

> A Vue.js project

## Build Setup

``` bash
# install dependencies
yarn install

(ps:如果报错node-sass相关问题，to run: SASS_BINARY_SITE=http://npm.taobao.org/mirrors/node-sass npm rebuild node-sass)


# serve with hot reload at localhost:8989
yarn start


# build for production with minification
yarn build

# build for production and view the bundle analyzer report
yarn build --report
```

## Structure

- ./build/webpack.base.js -Webpack 配置文件（基础配置文件）
- ./build/webpack.dev.js - Webpack 配置文件 (开发环境)
- ./build/webpack.prod.js - Webpack 配置文件 (生产环境)
- ./dist/ - 打包后的代码
- ./docs/ - 相关文档
- ./src - 项目源码
- ./src/index.html - Webpack 打包用模板
- ./src/assets/ - 资源文件
- ./src/assets/styles - 全局的 CSS 文件
- ./src/assets/images - 全局的图片资源
- ./src/main.js - 应用入口
- ./src/App.vue - 根组件
- ./src/shims.d.ts - typescript配置
- ./src/views/ - 视图
- ./src/router/ - 路由
- ./src/vendors/ - 依赖库入口
- ./src/store/ - 状态管理
- ./src/utils/ - 抽象工具
- ./src/api/ - API请求
- ./src/components/ - 自定义公共组件

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).
