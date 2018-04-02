/*
* @Author: jensen
* @Date:   2018-04-02 15:36:05
* @Last Modified by:   jensen
* @Last Modified time: 2018-04-02 18:22:27
*/

import Vue from 'vue'

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
  CarouselItem,
} from 'element-ui'
import CollapseTransition from 'element-ui/lib/transitions/collapse-transition'

Vue.$message = Vue.prototype.$message = Message

// register
Vue.useAll(
  Row, 
  Col, 
  Input, 
  Button, 
  Form, 
  FormItem, 
  Checkbox, 
  Radio, 
  RadioGroup,
  Select, 
  Option, 
  Tree, 
  Dialog, 
  Progress, 
  Tag, 
  Pagination, 
  Dropdown, 
  DropdownMenu, 
  DropdownItem,Tabs, 
  TabPane, Breadcrumb, 
  BreadcrumbItem, Card, 
  Popover, 
  DatePicker, 
  Upload,Table, 
  TableColumn, 
  Scrollbar, 
  Cascader, 
  Switch, 
  Carousel, 
  CarouselItem, 
  CollapseTransition
)
