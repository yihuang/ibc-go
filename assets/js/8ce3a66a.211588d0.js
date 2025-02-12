"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[76474],{41702:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>r,contentTitle:()=>o,default:()=>h,frontMatter:()=>c,metadata:()=>a,toc:()=>d});var n=t(85893),i=t(11151);const c={title:"Messages",sidebar_label:"Messages",sidebar_position:4,slug:"/ibc/light-clients/wasm/messages"},o="Messages",a={id:"light-clients/wasm/messages",title:"Messages",description:"MsgStoreCode",source:"@site/versioned_docs/version-v9.0.x/03-light-clients/04-wasm/04-messages.md",sourceDirName:"03-light-clients/04-wasm",slug:"/ibc/light-clients/wasm/messages",permalink:"/v9/ibc/light-clients/wasm/messages",draft:!1,unlisted:!1,tags:[],version:"v9.0.x",sidebarPosition:4,frontMatter:{title:"Messages",sidebar_label:"Messages",sidebar_position:4,slug:"/ibc/light-clients/wasm/messages"},sidebar:"defaultSidebar",previous:{title:"Integration",permalink:"/v9/ibc/light-clients/wasm/integration"},next:{title:"Governance",permalink:"/v9/ibc/light-clients/wasm/governance"}},r={},d=[{value:"<code>MsgStoreCode</code>",id:"msgstorecode",level:2},{value:"<code>MsgMigrateContract</code>",id:"msgmigratecontract",level:2},{value:"<code>MsgRemoveChecksum</code>",id:"msgremovechecksum",level:2}];function l(e){const s={a:"a",code:"code",h1:"h1",h2:"h2",li:"li",p:"p",pre:"pre",ul:"ul",...(0,i.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(s.h1,{id:"messages",children:"Messages"}),"\n",(0,n.jsx)(s.h2,{id:"msgstorecode",children:(0,n.jsx)(s.code,{children:"MsgStoreCode"})}),"\n",(0,n.jsxs)(s.p,{children:["Uploading the Wasm light client contract to the Wasm VM storage is achieved by means of ",(0,n.jsx)(s.code,{children:"MsgStoreCode"}),":"]}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{className:"language-go",children:"type MsgStoreCode struct {\n  // signer address\n  Signer string\n  // wasm byte code of light client contract. It can be raw or gzip compressed\n  WasmByteCode []byte\n}\n"})}),"\n",(0,n.jsx)(s.p,{children:"This message is expected to fail if:"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"Signer"})," is an invalid Bech32 address, or it does not match the designated authority address."]}),"\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"WasmByteCode"})," is empty or it exceeds the maximum size, currently set to 3MB."]}),"\n"]}),"\n",(0,n.jsxs)(s.p,{children:["Only light client contracts stored using ",(0,n.jsx)(s.code,{children:"MsgStoreCode"})," are allowed to be instantiated. An attempt to create a light client from contracts uploaded via other means (e.g. through ",(0,n.jsx)(s.code,{children:"x/wasm"})," if the module shares the same Wasm VM instance with 08-wasm) will fail. Due to the idempotent nature of the Wasm VM's ",(0,n.jsx)(s.code,{children:"StoreCode"})," function, it is possible to store the same byte code multiple times."]}),"\n",(0,n.jsxs)(s.p,{children:["When execution of ",(0,n.jsx)(s.code,{children:"MsgStoreCode"})," succeeds, the checksum of the contract (i.e. the sha256 hash of the contract's byte code) is stored in an allow list. When a relayer submits ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/proto/ibc/core/client/v1/tx.proto#L25-L37",children:(0,n.jsx)(s.code,{children:"MsgCreateClient"})})," with 08-wasm's ",(0,n.jsx)(s.code,{children:"ClientState"}),", the client state includes the checksum of the Wasm byte code that should be called. Then 02-client calls ",(0,n.jsxs)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/06fd8eb5ee1697e3b43be7528a6e42f5e4a4613c/modules/core/02-client/keeper/client.go#L40",children:["08-wasm's implementation of ",(0,n.jsx)(s.code,{children:"Initialize"})," function"]})," (which is an interface function part of ",(0,n.jsx)(s.code,{children:"LightClientModule"}),"), and it will check that the checksum in the client state matches one of the checksums in the allow list. If a match is found, the light client is initialized; otherwise, the transaction is aborted."]}),"\n",(0,n.jsx)(s.h2,{id:"msgmigratecontract",children:(0,n.jsx)(s.code,{children:"MsgMigrateContract"})}),"\n",(0,n.jsxs)(s.p,{children:["Migrating a contract to a new Wasm byte code is achieved by means of ",(0,n.jsx)(s.code,{children:"MsgMigrateContract"}),":"]}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{className:"language-go",children:"type MsgMigrateContract struct {\n  // signer address\n  Signer string\n  // the client id of the contract\n  ClientId string\n  // the SHA-256 hash of the new wasm byte code for the contract\n  Checksum []byte\n  // the json-encoded migrate msg to be passed to the contract on migration\n  Msg []byte\n}\n"})}),"\n",(0,n.jsx)(s.p,{children:"This message is expected to fail if:"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"Signer"})," is an invalid Bech32 address, or it does not match the designated authority address."]}),"\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"ClientId"})," is not a valid identifier prefixed by ",(0,n.jsx)(s.code,{children:"08-wasm"}),"."]}),"\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"Checksum"})," is not exactly 32 bytes long or it is not found in the list of allowed checksums (a new checksum is added to the list when executing ",(0,n.jsx)(s.code,{children:"MsgStoreCode"}),"), or it matches the current checksum of the contract."]}),"\n"]}),"\n",(0,n.jsx)(s.p,{children:"When a Wasm light client contract is migrated to a new Wasm byte code the checksum for the contract will be updated with the new checksum."}),"\n",(0,n.jsx)(s.h2,{id:"msgremovechecksum",children:(0,n.jsx)(s.code,{children:"MsgRemoveChecksum"})}),"\n",(0,n.jsxs)(s.p,{children:["Removing a checksum from the list of allowed checksums is achieved by means of ",(0,n.jsx)(s.code,{children:"MsgRemoveChecksum"}),":"]}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{className:"language-go",children:"type MsgRemoveChecksum struct {\n  // signer address\n  Signer string\n  // Wasm byte code checksum to be removed from the store\n  Checksum []byte\n}\n"})}),"\n",(0,n.jsx)(s.p,{children:"This message is expected to fail if:"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"Signer"})," is an invalid Bech32 address, or it does not match the designated authority address."]}),"\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"Checksum"})," is not exactly 32 bytes long or it is not found in the list of allowed checksums (a new checksum is added to the list when executing ",(0,n.jsx)(s.code,{children:"MsgStoreCode"}),")."]}),"\n"]}),"\n",(0,n.jsxs)(s.p,{children:["When a checksum is removed from the list of allowed checksums, then the corresponding Wasm byte code will not be available for instantiation in ",(0,n.jsxs)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/core/02-client/keeper/client.go#L36",children:["08-wasm's implementation of ",(0,n.jsx)(s.code,{children:"Initialize"})," function"]}),"."]})]})}function h(e={}){const{wrapper:s}={...(0,i.a)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(l,{...e})}):l(e)}},11151:(e,s,t)=>{t.d(s,{Z:()=>a,a:()=>o});var n=t(67294);const i={},c=n.createContext(i);function o(e){const s=n.useContext(c);return n.useMemo((function(){return"function"==typeof e?e(s):{...s,...e}}),[s,e])}function a(e){let s;return s=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:o(e.components),n.createElement(c.Provider,{value:s},e.children)}}}]);