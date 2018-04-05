/*
* @Author: jensen
* @Date:   2018-04-05 13:38:25
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-05 13:43:20
*/

import Vue from 'vue';

/**
 * import element-ui COMMON components here
 */
import {
  Row,
  Col,
  Form,
  FormItem,
  Checkbox,
  Radio,
  RadioGroup,
  Input,
  Button,
  Cascader,
  Message,
  Upload,
  Select,
  Option,
  DatePicker,
  Tree,
  Dialog,
  Progress,
  Tag,
  Pagination,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Tabs,
  TabPane,
  Breadcrumb,
  BreadcrumbItem,
  Card,
  Popover,
  Table,
  TableColumn,
  Scrollbar,
  Loading,
  Switch,
  Carousel,
  CarouselItem
} from 'element-ui';
import CollapseTransition from 'element-ui/lib/transitions/collapse-transition';

Vue.components(Row, Col, Input, Button, Form, FormItem, Checkbox, Radio, RadioGroup,
  Select, Option, Tree, Dialog, Progress, Tag, Pagination, Dropdown, DropdownMenu, DropdownItem,
  Tabs, TabPane, Breadcrumb, BreadcrumbItem, Card, Popover, DatePicker, Upload,
  Table, TableColumn, Scrollbar, Cascader, Switch, Carousel, CarouselItem, CollapseTransition);


Vue.prototype.$message = Message;

Vue.use(Loading.directive);
Vue.prototype.$loading = Loading.service;


