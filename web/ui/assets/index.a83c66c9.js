var p=(t,i,a)=>new Promise((s,r)=>{var c=e=>{try{u(a.next(e))}catch(n){r(n)}},l=e=>{try{u(a.throw(e))}catch(n){r(n)}},u=e=>e.done?s(e.value):Promise.resolve(e.value).then(c,l);u((a=a.apply(t,i)).next())});import{B as b,u as D}from"./useTable.35e7adb4.js";import{T as j}from"./useForm.d0a5ee18.js";import{b as B}from"./index.24728625.js";import{D as g,c as h,s as E}from"./DatasourceModal.09cfddaa.js";import{s as x,i as F,d as _}from"./index.e064266a.js";import{_ as T,h as M}from"./index.c0a3bc85.js";import{A,a0 as d,B as w,D as S,k as m,a5 as f,ad as k}from"./vendor.52d192f7.js";import"./index.1c6370f0.js";import"./index.cd360e12.js";/* empty css               *//* empty css               */import"./useWindowSizeFn.34f605a6.js";import"./useContentViewHeight.26794800.js";/* empty css               *//* empty css                *//* empty css               *//* empty css               */import"./useSortable.96f14e96.js";/* empty css               *//* empty css              *//* empty css               */import"./index.1c8e4638.js";import"./download.0ca988cf.js";const R=A({name:"DatasourceManagement",components:{BasicTable:b,DatasourceModal:g,TableAction:j},setup(){const[t,{openModal:i}]=B(),{createMessage:a,createErrorModal:s}=M(),[r,{reload:c}]=D({title:"",api:x,columns:h,formConfig:{labelWidth:120,schemas:E},useSearchForm:!0,showTableSetting:!0,bordered:!0,striped:!0,pagination:!1,showIndexColumn:!1,actionColumn:{width:130,title:"\u64CD\u4F5C",dataIndex:"action",slots:{customRender:"action"},fixed:void 0}});function l(){i(!0,{isUpdate:!1})}function u(o){return p(this,null,function*(){if(yield F(o.id)){s({title:"\u9519\u8BEF\u63D0\u793A",content:"\u8BE5\u6570\u636E\u6E90\u6B63\u5728\u8FD0\u884C\u4E2D\uFF0C\u65E0\u6CD5\u7F16\u8F91\uFF1B\u8BF7\u5148\u5728'\u8FD0\u884C\u7BA1\u7406'\u6A21\u5757\u4E2D\u505C\u6B62"});return}i(!0,{record:o,isUpdate:!0})})}function e(o){return p(this,null,function*(){if(yield F(o.id)){s({title:"\u9519\u8BEF\u63D0\u793A",content:"\u8BE5\u6570\u636E\u6E90\u6B63\u5728\u8FD0\u884C\u4E2D\uFF0C\u65E0\u6CD5\u5220\u9664\uFF1B\u8BF7\u5148\u5728'\u8FD0\u884C\u7BA1\u7406'\u6A21\u5757\u4E2D\u505C\u6B62"});return}_(o.id).then(()=>{n()})})}function n(){c()}return{registerTable:r,registerModal:t,createMessage:a,createErrorModal:s,handleCreate:l,handleEdit:u,handleDelete:e,handleSuccess:n}}}),v=k(" \u65B0\u589E ");function y(t,i,a,s,r,c){const l=d("a-button"),u=d("TableAction"),e=d("BasicTable"),n=d("DatasourceModal");return w(),S("div",null,[m(e,{size:"small",onRegister:t.registerTable},{toolbar:f(()=>[m(l,{type:"primary",onClick:t.handleCreate},{default:f(()=>[v]),_:1},8,["onClick"])]),action:f(({record:o})=>[m(u,{actions:[{label:"\u7F16\u8F91",icon:"clarity:note-edit-line",onClick:t.handleEdit.bind(null,o)},{label:"\u5220\u9664",icon:"ant-design:delete-outlined",color:"error",popConfirm:{title:"\u662F\u5426\u786E\u8BA4\u5220\u9664",confirm:t.handleDelete.bind(null,o)}}]},null,8,["actions"])]),_:1},8,["onRegister"]),m(n,{onRegister:t.registerModal,onSuccess:t.handleSuccess},null,8,["onRegister","onSuccess"])])}var ue=T(R,[["render",y]]);export{ue as default};
