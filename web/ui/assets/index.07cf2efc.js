import{B as m,u as p}from"./useTable.35e7adb4.js";import{T as x}from"./useForm.d0a5ee18.js";import{w as a,bh as r,x as f,A as j,a0 as d,B,D as b,k as u,a5 as h}from"./vendor.52d192f7.js";/* empty css               */import{_ as T,ag as g,ah as C}from"./index.c0a3bc85.js";import"./index.1c6370f0.js";import"./index.cd360e12.js";/* empty css               *//* empty css               */import"./useWindowSizeFn.34f605a6.js";import"./useContentViewHeight.26794800.js";/* empty css               *//* empty css                */import"./index.24728625.js";/* empty css               */import"./useSortable.96f14e96.js";/* empty css               *//* empty css              *//* empty css               */import"./index.1c8e4638.js";import"./download.0ca988cf.js";const _=[{title:"\u8282\u70B9",dataIndex:"addr"},{title:"\u7C7B\u578B",dataIndex:"isLeader",width:200,customRender:({record:e})=>{const t=e.isLeader,o=t?"green":"blue",n=t?"\u4E3B\u8282\u70B9":"\u4ECE\u8282\u70B9";return a(r,{color:o},()=>n)}},{title:"\u72B6\u6001",dataIndex:"port",width:100,customRender:({record:e})=>{const t=e.deadline,o=t?"red":"green",n=t?"\u79BB\u7EBF":"\u6B63\u5E38";return a(r,{color:o},()=>n)}},{title:"\u6700\u540E\u6D3B\u8DC3\u65F6\u95F4",dataIndex:"lastActiveTime",width:250,customRender:({record:e})=>e.lastActiveTime==0?"":f.unix(e.lastActiveTime).format("YYYY-MM-DD HH:mm:ss")}],w=j({name:"ClusterManagement",components:{BasicTable:m,TableAction:x},setup(){const[e,{reload:t}]=p({title:"",api:g,columns:_,useSearchForm:!1,showTableSetting:!0,bordered:!0,striped:!0,pagination:!1,showIndexColumn:!1,actionColumn:{width:130,title:"\u64CD\u4F5C",dataIndex:"action",slots:{customRender:"action"},fixed:void 0}});function o(i){C(i.addr).then(()=>{n()})}function n(){t()}return{registerTable:e,handleDelete:o,handleSuccess:n}}});function v(e,t,o,n,i,A){const l=d("TableAction"),c=d("BasicTable");return B(),b("div",null,[u(c,{size:"small",onRegister:e.registerTable},{action:h(({record:s})=>[u(l,{actions:[{label:"\u4E0B\u7EBF",icon:"ant-design:delete-outlined",color:"error",onClick:e.handleDelete.bind(null,s),ifShow:E=>s.deadline}]},null,8,["actions"])]),_:1},8,["onRegister"])])}var P=T(w,[["render",v]]);export{P as default};
