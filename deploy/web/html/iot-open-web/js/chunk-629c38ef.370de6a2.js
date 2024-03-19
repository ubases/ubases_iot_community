(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-629c38ef","chunk-2d21711c","chunk-2d21711c"],{4966:function(t,e,r){"use strict";r("b0c0");var n=function(){var t=this,e=t._self._c;return e("div",{staticClass:"product-details flex y-axis-center"},[e("section",{staticClass:"image-wrap"},[e("img",{staticClass:"product-image",attrs:{src:t.productDetails.imageUrl||""}})]),e("section",{staticClass:"details"},[e("p",{staticClass:"product-name",domProps:{textContent:t._s(t.productDetails.name)}}),e("section",[e("a-row",{attrs:{gutter:40,type:"flex"}},t._l(t.detailList,(function(r){return e("a-col",{key:r},[t._v(" "+t._s(t.detailOrigin[r])+"："+t._s(t.productDetails[r])+" ")])})),1)],1)]),t.exchangeBtn?e("section",{staticClass:"exchange-product-wrap"},[e("a-button",{on:{click:function(e){return e.stopPropagation(),t.exchangeProduct.apply(null,arguments)}}},[e("exchange-icon"),t._v(t._s(t.$t("productChange.change.product")))],1),t.showProductSearch?e("a-select",{staticClass:"product-search",attrs:{defaultOpen:"","show-search":"",placeholder:t.$t("productChange.placeholder.product.search")},on:{change:t.handleProductChange}},t._l(t.productList,(function(r){return e("a-select-option",{key:r.id,attrs:{value:r.name+"_"+r.id}},[e("section",{staticClass:"flex y-axis-center"},[e("h4",[t._v(t._s(r.name))]),e("div",{staticStyle:{"margin-left":"10px",height:"22px","font-size":"14px","line-height":"20px",padding:"0 10px",border:"1px solid #DAE6F8","border-radius":"14px"}},[t._v(t._s(r.powerConsumeTypeDesc))])]),e("section",{staticClass:"flex y-axis-center x-space-between"},[e("span",[t._v(t._s(t.$t("productChange.productKey"))+"："+t._s(r.productKey))]),e("span",[t._v(t._s(t.$t("productChange.productTypeName"))+"："+t._s(r.productTypeName))]),e("span",[t._v(t._s(t.$t("productChange.networkTypeDesc"))+"："+t._s(r.networkTypeDesc))])])])})),1):t._e()],1):t._e()])},o=[],u=r("c7eb"),c=r("5530"),a=r("1da1"),p=(r("4de4"),r("d3b7"),r("04b3")),d=r("c4c8"),i={props:{exchangeBtn:{type:Boolean,default:!0},productId:{type:String,default:""},detailList:{type:Array,default:function(){return["productTypeName","productKey","networkTypeDesc","powerConsumeTypeDesc"]}},param:{type:Object,default:function(){}}},components:{ExchangeIcon:p["q"]},data:function(){return{productDetails:{},productList:[],showProductSearch:!1,detailOrigin:{productTypeName:this.$t("productChange.detailOrigin.productTypeName"),productKey:this.$t("productChange.detailOrigin.productKey"),networkTypeDesc:this.$t("productChange.detailOrigin.networkTypeDesc"),powerConsumeTypeDesc:this.$t("productChange.detailOrigin.powerConsumeTypeDesc")}}},created:function(){this.queryProductList()},methods:{queryProductList:function(){var t=this;return Object(a["a"])(Object(u["a"])().mark((function e(){var r,n,o,a;return Object(u["a"])().wrap((function(e){while(1)switch(e.prev=e.next){case 0:return o={page:0},t.param&&(o=Object(c["a"])(Object(c["a"])({},o),t.param)),e.next=4,Object(d["A"])(o);case 4:if(a=e.sent,t.$emit("isData",0==a.code&&a.data.list&&a.data.list.length>0),0===a.code){e.next=8;break}return e.abrupt("return");case 8:t.productList=null!==(r=null===(n=a.data)||void 0===n?void 0:n.list)&&void 0!==r?r:[],t.productId?t.productDetails=t.productList.filter((function(e){return e.id==t.productId})).pop()||{}:t.productDetails=t.productList.length>0?t.productList[0]:{},t.$emit("dataChange",t.productDetails);case 11:case"end":return e.stop()}}),e)})))()},exchangeProduct:function(){this.showProductSearch=!this.showProductSearch},handleProductChange:function(t){this.showProductSearch=!this.showProductSearch;var e=t.split("_").pop();this.productDetails=this.productList.filter((function(t){return t.id==e})).pop(),this.$emit("dataChange",this.productDetails)}}},s=i,l=(r("578c"),r("2877")),f=Object(l["a"])(s,n,o,!1,null,"2209c8d0",null);e["a"]=f.exports},"501e":function(t,e,r){"use strict";var n=function(){var t=this,e=t._self._c;return e("section",{staticClass:"no-data-wrap"},[e("img",{attrs:{src:r("8c0b")}}),e("h3",[t._v(t._s(t.$t("public.no.data")))]),e("p",{staticClass:"details",domProps:{textContent:t._s(t.text)}}),e("a-button",{staticClass:"grean-button",attrs:{type:"primary"},on:{click:t.add}},[e("icon-font",{style:{fontSize:"18px"},attrs:{type:"icon-add"}}),t._v(" "+t._s(t.buttonText)+" ")],1)],1)},o=[],u={name:"NoData",props:{text:{type:String,default:""},buttonText:{type:String,default:""},path:{type:String,default:""}},methods:{add:function(){this.$router.push({path:this.path})}}},c=u,a=(r("7210"),r("2877")),p=Object(a["a"])(c,n,o,!1,null,"fedf93fc",null);e["a"]=p.exports},"578c":function(t,e,r){"use strict";r("ec88")},7210:function(t,e,r){"use strict";r("8b64")},"8b64":function(t,e,r){},c4c8:function(t,e,r){"use strict";r.d(e,"A",(function(){return o})),r.d(e,"z",(function(){return u})),r.d(e,"j",(function(){return c})),r.d(e,"y",(function(){return a})),r.d(e,"g",(function(){return p})),r.d(e,"m",(function(){return d})),r.d(e,"F",(function(){return i})),r.d(e,"t",(function(){return s})),r.d(e,"C",(function(){return l})),r.d(e,"c",(function(){return f})),r.d(e,"k",(function(){return b})),r.d(e,"a",(function(){return m})),r.d(e,"h",(function(){return h})),r.d(e,"v",(function(){return w})),r.d(e,"Q",(function(){return v})),r.d(e,"n",(function(){return g})),r.d(e,"e",(function(){return O})),r.d(e,"q",(function(){return j})),r.d(e,"p",(function(){return y})),r.d(e,"P",(function(){return C})),r.d(e,"K",(function(){return x})),r.d(e,"x",(function(){return _})),r.d(e,"R",(function(){return D})),r.d(e,"w",(function(){return T})),r.d(e,"r",(function(){return k})),r.d(e,"N",(function(){return L})),r.d(e,"d",(function(){return P})),r.d(e,"o",(function(){return S})),r.d(e,"f",(function(){return $})),r.d(e,"M",(function(){return F})),r.d(e,"E",(function(){return N})),r.d(e,"D",(function(){return G})),r.d(e,"U",(function(){return q})),r.d(e,"H",(function(){return K})),r.d(e,"T",(function(){return E})),r.d(e,"G",(function(){return I})),r.d(e,"O",(function(){return A})),r.d(e,"B",(function(){return B})),r.d(e,"I",(function(){return z})),r.d(e,"J",(function(){return J})),r.d(e,"L",(function(){return M})),r.d(e,"S",(function(){return R})),r.d(e,"s",(function(){return V})),r.d(e,"b",(function(){return U})),r.d(e,"l",(function(){return H})),r.d(e,"i",(function(){return Q})),r.d(e,"u",(function(){return W}));var n=r("b775");function o(t){return Object(n["b"])({url:"/v1/platform/web/open/product/list",method:"post",data:t})}function u(t){return Object(n["b"])({url:"/v1/platform/web/open/product/detail/"+t,method:"get"})}function c(t){return Object(n["b"])({url:"/v1/platform/web/open/product/delete",method:"post",data:t})}function a(){return Object(n["b"])({url:"/v1/platform/web/open/productType/get",method:"get"})}function p(t){return Object(n["b"])({url:"/v1/platform/web/open/product/save",method:"post",data:t})}function d(t){return Object(n["b"])({url:"/v1/platform/web/open/product/edit",method:"post",data:t})}function i(t){return Object(n["b"])({url:"/v1/platform/web/open/product/functions",method:"get",params:t})}function s(t){return Object(n["b"])({url:"/v1/platform/web/open/product/funcList",method:"get",params:t})}function l(t){return Object(n["b"])({url:"/v1/platform/web/open/product/standardFuncList",method:"get",params:t})}function f(t){return Object(n["b"])({url:"/v1/platform/web/open/product/addStandardFunc",method:"post",data:t})}function b(t){return Object(n["b"])({url:"/v1/platform/web/open/product/editFunc",method:"post",data:t})}function m(t){return Object(n["b"])({url:"/v1/platform/web/open/product/addFunc",method:"post",data:t})}function h(t){return Object(n["b"])({url:"/v1/platform/web/open/product/deleteFunc",method:"post",data:t})}function w(t){return Object(n["b"])({url:"/v1/platform/web/open/product/queryModules",method:"get",params:t})}function v(t){return Object(n["b"])({url:"/v1/platform/web/open/product/selectModule",method:"post",data:t})}function g(t){return Object(n["b"])({url:"/v1/platform/web/open/firmware/changeVersionList",method:"get",params:t})}function O(t){return Object(n["b"])({url:"/v1/platform/web/open/product/changeVersionSubmit",method:"post",data:t})}function j(t){return Object(n["b"])({url:"/v1/platform/web/open/product/queryCustomFirmware",method:"get",params:t})}function y(t){return Object(n["b"])({url:"/v1/platform/web/open/firmware/changeCustomVersionList",method:"get",params:t})}function C(t){return Object(n["b"])({url:"/v1/platform/web/open/product/selectCustomerFirmware",method:"post",data:t})}function x(t){return Object(n["b"])({url:"/v1/platform/web/open/product/removeCustomerFirmware",method:"post",data:t})}function _(t){return Object(n["b"])({url:"/v1/platform/web/open/product/controlPanelList",method:"get",params:t})}function D(t){return Object(n["b"])({url:"/v1/platform/web/open/product/selectControlPanel",method:"post",data:t})}function T(t){return Object(n["b"])({url:"/v1/platform/web/open/product/getNetworkGuide",method:"get",params:t})}function k(t){return Object(n["b"])({url:"/v1/platform/web/open/product/getDefaultNetworkGuide",method:"get",params:t})}function L(t){return Object(n["b"])({url:"/v1/platform/web/open/product/saveNetworkGuide",method:"post",data:t})}function P(t){return Object(n["b"])({url:"/v1/platform/web/open/product/changeNetworkGuide",method:"post",data:t})}function S(t){return Object(n["b"])({url:"/v1/platform/web/open/product/completeDevelopDetailed",method:"get",params:t})}function $(t){return Object(n["b"])({url:"/v1/platform/web/open/product/completeDevelop",method:"post",data:t})}function F(t){return Object(n["b"])({url:"/v1/platform/web/open/product/returnDevelop",method:"post",data:t})}function N(t){return Object(n["b"])({url:"/v1/platform/web/template/testCaseTpl/download",method:"get",params:t})}function G(t){return Object(n["b"])({url:"/v1/platform/web/open/product/getTestReportFile",method:"get",params:t})}function q(t){return Object(n["b"])({url:"/v1/platform/web/open/product/uploadTestReport",method:"post",timeout:3e4,data:t})}function K(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/list",method:"post",data:t})}function E(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/publish",method:"post",data:t})}function I(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/detail",method:"post",data:t})}function A(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/save",method:"post",data:t})}function B(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/publish/record",method:"get",params:t})}function z(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/getDoc",method:"get",params:t})}function J(t){return Object(n["b"])({url:"/v1/platform/web/open/product/voice/unitList",method:"post",data:t})}function M(t){return Object(n["b"])({url:"/v1/platform/web/open/product/resetStandardFunc",method:"GET",params:t})}function R(t){return Object(n["b"])({url:"/v1/platform/web/open/product/setSceneFunc",method:"post",data:t})}function V(t){return Object(n["b"])({url:"/v1/platform/web/open/product/funcList",method:"get",params:t})}function U(t){return Object(n["b"])({url:"/v1/platform/web/open/manual/add",method:"post",data:t})}function H(t){return Object(n["b"])({url:"/v1/platform/web/open/manual/edit",method:"post",data:t})}function Q(t){return Object(n["b"])({url:"/v1/platform/web/open/manual/del",method:"post",data:t})}function W(t){return Object(n["b"])({url:"/v1/platform/web/open/manual/detail",method:"GET",params:t})}},ec88:function(t,e,r){}}]);