(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7e377606"],{"5e89":function(t,e,a){"use strict";a("b0c0");var n=function(){var t=this,e=t._self._c;return e("section",["table"===t.type?e("a-row",{staticClass:"thread"},t._l(t.columnList,(function(a){return e("a-col",{key:a.name,attrs:{span:a.span}},[t._v(" "+t._s(a.name)+" ")])})),1):t._e(),e("section",t._l(t.dataList,(function(a,n){return e("a-row",{key:n,class:["card",{active:1===a.active},{disabled:a.disabled}],attrs:{type:"flex",align:"middle"}},t._l(t.columnList,(function(n){return e("a-col",{key:n.slot,attrs:{span:n.span}},[t._t(n.slot,(function(){return[t._v(t._s(a[n.slot]))]}),{record:a})],2)})),1)})),1)],1)},s=[],r={name:"TableCard",props:{columnList:{type:Array,default:function(){return[]}},dataList:{type:Array,default:function(){return[]}},type:{type:String,default:"table"}},data:function(){return{}},methods:{}},c=r,i=(a("9dbd"),a("2877")),o=Object(i["a"])(c,n,s,!1,null,"7f26ba9c",null);e["a"]=o.exports},"6bb2":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t._self._c;return e("table-card",t._b({scopedSlots:t._u([{key:"moduleName",fn:function(a){var n=a.record;return[e("section",{staticClass:"flex y-axis-center"},[e("section",{staticClass:"image-wrap"},[e("img",{staticClass:"module-image",attrs:{src:n.imgUrl}})]),e("section",[e("p",{staticClass:"module-name",domProps:{textContent:t._s(n.moduleName)}}),e("a",{staticClass:"check-doc",attrs:{target:"_blank",href:n.fileUrl}},[t._v(t._s(t.$t("setting.hardware.develop.check.doc"))+" >")])])])]}},{key:"type",fn:function(a){var n=a.record;return[e("p",[t._v(t._s(t.$DictName("firmware_type",n.type)))])]}},{key:"action",fn:function(a){var n=a.record;return[0!==n.active||n.disabled?t._e():e("a-button",{attrs:{type:"link"},on:{click:function(e){return t.handleSelect(n)}}},[t._v(t._s(t.$t("setting.hardware.develop.select")))]),1===n.active?e("a-button",{attrs:{type:"link",disabled:""}},[t._v(t._s(t.$t("setting.hardware.develop.selected")))]):t._e(),n.associated?e("a-button",{attrs:{type:"link",disabled:""}},[t._v(t._s(t.$t("setting.hardware.develop.associated")))]):t._e()]}}])},"table-card",t.$attrs,!1))},s=[],r=a("c7eb"),c=a("1da1"),i=(a("a9e3"),a("5e89")),o={name:"HardwareTableCard",inheritAttrs:!0,props:{productId:{type:String,default:""},status:{type:Number,default:0}},components:{TableCard:i["a"]},data:function(){return{}},methods:{handleSelect:function(t){var e=this;return Object(c["a"])(Object(r["a"])().mark((function a(){return Object(r["a"])().wrap((function(a){while(1)switch(a.prev=a.next){case 0:e.$emit("handleSelect",t);case 1:case"end":return a.stop()}}),a)})))()},disassociate:function(t){this.$emit("disassociate",t)}}},d=o,l=(a("fd6f"),a("2877")),u=Object(l["a"])(d,n,s,!1,null,"3dede7b4",null);e["default"]=u.exports},"9dbd":function(t,e,a){"use strict";a("a749")},a749:function(t,e,a){},c6d4:function(t,e,a){},fd6f:function(t,e,a){"use strict";a("c6d4")}}]);