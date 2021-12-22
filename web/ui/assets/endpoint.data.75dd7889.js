import{aH as a,aI as o}from"./index.c0a3bc85.js";import{w as F,bh as r}from"./vendor.52d192f7.js";/* empty css               */var s;(function(u){u.EndpointUrl="/endpoints"})(s||(s={}));const A=u=>a.post({url:s.EndpointUrl,params:u},{errorMessageMode:"modal"}),i=u=>a.put({url:s.EndpointUrl,params:u},{errorMessageMode:"modal"}),E=u=>a.delete({url:o(s.EndpointUrl,u)},{errorMessageMode:"modal"}),b=u=>a.get({url:s.EndpointUrl,params:u}),g=u=>a.get({url:s.EndpointUrl,params:u}),f=u=>a.post({url:o(s.EndpointUrl,"test-link"),params:u},{errorMessageMode:"modal"}),B=u=>a.get({url:o(s.EndpointUrl,u,"is-running")},{errorMessageMode:"modal"});function p(u){let e="";switch(u){case 1:{e="Redis";break}case 2:{e="MongoDB";break}case 3:{e="Elasticsearch";break}case 4:{e="ClickHouse";break}case 5:{e="RocketMQ";break}case 6:{e="Kafka";break}case 7:{e="RabbitMQ";break}case 8:{e="HTTP\u63A5\u53E3";break}case 9:{e="GRPC\u63A5\u53E3";break}}return e}const P=[{title:"\u540D\u79F0",dataIndex:"name"},{title:"\u7C7B\u578B",dataIndex:"type",width:160,customRender:({record:u})=>p(u.type)},{title:"\u5730\u5740",dataIndex:"addresses",width:300},{title:"\u72B6\u6001",dataIndex:"status",width:80,customRender:({record:u})=>{const l=~~u.status==0,n=l?"green":"red",t=l?"\u542F\u7528":"\u505C\u7528";return F(r,{color:n},()=>t)}}],C=[{field:"name",label:"\u540D\u79F0",component:"Input",colProps:{span:8}},{field:"host",label:"\u4E3B\u673A",component:"Input",colProps:{span:8}}],h=[{field:"name",label:"\u540D\u79F0",component:"Input",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"addresses",label:"\u5730\u5740",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3\uFF0C\u5982\uFF1A127.0.0.1:6379","\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:6379,127.0.0.2:6379"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:6379\uFF0C\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:6379,127.0.0.2:6379"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u5730\u5740"},{max:420,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"deployType",label:"\u90E8\u7F72\u6A21\u5F0F",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u5355\u673A",value:0},{label:"\u96C6\u7FA4",value:1}]}},{field:"groupType",label:"\u96C6\u7FA4\u6A21\u5F0F",required:!0,component:"Select",ifShow:({values:u})=>u.deployType===1,colProps:{span:12},componentProps:{options:[{label:"sentinel",value:0,key:"sentinel"},{label:"cluster",value:1,key:"cluster"}]}},{field:"masterName",label:"Master Name",required:!0,component:"Input",colProps:{span:12},ifShow:({values:u})=>u.deployType===1&&u.groupType===0},{field:"database",label:"database",required:!0,component:"InputNumber",defaultValue:0,ifShow:({values:u})=>u.groupType!=2,componentProps:{max:16},colProps:{span:12}},{field:"password",label:"\u5BC6\u7801",component:"InputPassword",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:9}]}}],I=[{field:"name",label:"\u540D\u79F0",component:"Input",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"addresses",label:"\u5730\u5740",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3\uFF0C\u5982\uFF1A127.0.0.1:27017","\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:27017,127.0.0.2:27017"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:27017\uFF0C\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:27017,127.0.0.2:27017"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u5730\u5740"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"username",label:"\u7528\u6237\u540D",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"password",label:"\u5BC6\u7801",component:"InputPassword",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:1}]}}],D=[{field:"name",label:"\u540D\u79F0",component:"Input",colProps:{span:12},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"version",label:"\u7248\u672C",required:!0,component:"Select",colProps:{span:12},componentProps:{options:[{label:"6",value:6,key:"6"},{label:"7",value:7,key:"7"}]}},{field:"addresses",label:"\u5730\u5740",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3\uFF0C\u5982\uFF1A127.0.0.1:27017","\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:27017,127.0.0.2:27017"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:27017\uFF0C\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:27017,127.0.0.2:27017"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u5730\u5740"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"username",label:"\u7528\u6237\u540D",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"password",label:"\u5BC6\u7801",component:"InputPassword",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:1}]}}],x=[{field:"name",label:"\u540D\u79F0",component:"Input",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"addresses",label:"name_servers",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3\uFF0C\u5982\uFF1A127.0.0.1:9876","\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:9876,127.0.0.2:9876"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u547D\u540D\u670D\u52A1\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:9876\uFF0C\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:9876,127.0.0.2:9876"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u547D\u540D\u670D\u52A1\u5730\u5740"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"groupName",label:"Group Name",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"instanceName",label:"Instance Name",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"username",label:"access_key",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"password",label:"secret_key",component:"InputPassword",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A40"}]},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:1}]}}],k=[{field:"name",label:"\u540D\u79F0",component:"Input",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"addresses",label:"\u5730\u5740",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3\uFF0C\u5982\uFF1A127.0.0.1:9092","\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:9092,127.0.0.2:9092"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u547D\u540D\u670D\u52A1\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:9092\uFF0C\u591A\u4E2A\u5730\u5740\u7528\u82F1\u6587\u9017\u53F7\u5206\u5272,\u5982\uFF1A127.0.0.1:9876,127.0.0.2:9092"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u5730\u5740"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"username",label:"sasl_user",component:"Input",helpMessage:"kafka SASL_PLAINTEXT\u8BA4\u8BC1\u6A21\u5F0F \u7528\u6237\u540D",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"password",label:"sasl_password",component:"Input",helpMessage:"kafka SASL_PLAINTEXT\u8BA4\u8BC1\u6A21\u5F0F \u5BC6\u7801",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:1}]}}],M=[{field:"name",label:"\u540D\u79F0",component:"Input",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"addresses",label:"\u5730\u5740",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3\uFF0C\u5982\uFF1A127.0.0.1:9092"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u8FDE\u63A5\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:9092"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u5730\u5740"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"username",label:"\u7528\u6237\u540D",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"password",label:"\u5BC6\u7801",component:"Input",colProps:{span:12},rules:[{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:1}]}}],w=[{field:"name",label:"\u540D\u79F0",component:"Input",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A20"}]},{field:"addresses",label:"\u5730\u5740",component:"Input",helpMessage:["\u683C\u5F0F\uFF1AIP:\u7AEF\u53E3/*\uFF0C\u5982\uFF1A127.0.0.1:9092/*"],componentProps:{placeholder:"\u8BF7\u8F93\u5165\u8FDE\u63A5\u5730\u5740\uFF0C\u683C\u5F0F\uFF1A127.0.0.1:9092/*"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u5730\u5740"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A420"}]},{field:"authMode",label:"\u8BA4\u8BC1\u6A21\u5F0F",required:!0,component:"Select",helpMessage:["\u65E0\u9700\u8BA4\u8BC1\uFF1A\u53EF\u76F4\u63A5\u8BBF\u95EE\u63A5\u53E3",'HttpBasic\uFF1A\u5728\u8BF7\u6C42\u5934\u6DFB\u52A0\u5C5E\u6027Authorization\uFF0C\u503CBasic dGVzdDp0ZXN0\uFF0C\u5176\u4E2DdGVzdDp0ZXN0\u662F"username:password"\u7684Base64\u7F16\u7801',"JWT\uFF1A\u4F7F\u7528\u79D8\u94A5\u751F\u6210\u4E00\u4E2A\u6807\u51C6\u7684jwt\u4EE4\u724C"],componentProps:{options:[{label:"\u65E0\u9700\u8BA4\u8BC1",value:0,key:"0"},{label:"Http Basic",value:1,key:"1"},{label:"JWT",value:2,key:"2"}]}},{field:"username",label:"\u7528\u6237\u540D",component:"Input",colProps:{span:12},ifShow:({values:u})=>u.authMode==1,rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"password",label:"\u5BC6\u7801",component:"Input",colProps:{span:12},ifShow:({values:u})=>u.authMode==1,rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"jwtSecretKey",label:"JWT\u7B7E\u540D\u79D8\u94A5",component:"Input",colProps:{span:12},ifShow:({values:u})=>u.authMode==2,rules:[{required:!0,message:"\u8BF7\u8F93\u5165JWT\u7B7E\u540D\u79D8\u94A5"},{max:20,message:"\u6700\u5927\u957F\u5EA6\u4E3A255"}]},{field:"jwtExpire",label:"JWT\u7B7E\u540D\u6709\u6548\u671F(\u79D2)",component:"InputNumber",defaultValue:5,colProps:{span:12},componentProps:{min:5},ifShow:({values:u})=>u.authMode==2},{field:"status",label:"\u72B6\u6001",component:"RadioButtonGroup",defaultValue:0,colProps:{span:12},componentProps:{options:[{label:"\u542F\u7528",value:0},{label:"\u505C\u7528",value:1}]}}];export{C as a,A as b,P as c,E as d,D as e,x as f,M as g,w as h,B as i,g as j,k,p as l,I as m,i as p,h as r,b as s,f as t};
