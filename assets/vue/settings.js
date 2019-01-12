const Interface = {
  template: `<ul class="flx fcl fii pdd stg">
  <li><h3>Language</h3><button @click="getPrivateKey" class="btn">Get my Private Key</button></li>
  <li><h3>Language</h3><b-select placeholder="Select language"><option>English</option><option>Srpski</option><option>French</option></b-select></li>
  <li><h3>Coin denomination</h3><b-select placeholder="Coin denomination"><option>DUO</option><option>mDUO</option><option>uDuo</option></b-select></li>
  <li><h3>Fiat currency</h3><b-select placeholder="Fiat currency"><option>USD</option><option>EUR</option><option>RSD</option><option>BGN</option><option>CNY</option><option>CHF</option></b-select></li>
  <li><h3>Theme</h3><b-select placeholder="Theme"><option>Bright theme</option><option>Dark theme</option></b-select></li>
  <li><h3>Custom CSS</h3><b-input type="textarea" placeholder="css here"/></li>
  <li><h3>Custom CSS</h3><b-select placeholder="Start page"><option>Overview</option><option>History</option></b-select></li>   
  <li><h3>Minimise to tray</h3><b-switch v-model="showBooks">Minimise to tray</b-switch> </li>
  <li><h3>Custom CSS</h3><b-switch v-model="showBooks">Minimise instead of closing</b-switch></li>
  </ul>`
}
const Network = {
  template: `<ul class="flx fcl fii pdd stg">
  <li> <b-switch v-model="showBooks">Client TLS  (inbound)</b-switch> <b-switch v-model="showBooks">Client TLS  (inbound)</b-switch> </li>
  <li><h3>Peer to Peer Network</h3><b-input type="text" placeholder="host:port"/> <button class="btn">Change</button></li>
  <li><h3>RPC</h3><b-select  placeholder="RPC"><option>JSONRPC</option><option>gRPC</option><option>msgpack</option></b-select></li>
          <li><b-input type="text" placeholder="host:port"/> <button class="btn">Add</button></li>
  <li><h3>Language</h3>        type host:port <button>delete</button></li>
  <li><h3>Scrypt RPC</h3><b-select  placeholder="Scrypt RPC"><option>JSONRPC</option><option>gRPC</option><option>msgpack</option></b-select></li>
          <li><b-input type="text" placeholder="host:port"/> <button class="btn">Add</button></li>
  <li><h3>Language</h3>        host:port <button>delete</button></li>
  <li><h3>Legacy RPCe</h3><b-input type="text" placeholder="host:port"/> <button class="btn">Add</button></li>
  <li><h3>Language</h3>        host:port <button>delete</button></li>
  <li><h3>Language</h3><input type="radio">Add peers </li>
  <li><h3>Language</h3><b-input type="text" placeholder="host:port"/> <button class="btn">Add</button></li>
  <li><h3>Language</h3>        host:port status  <button>delete</button></li>
  <li><h3>Language</h3><b-input type="radio" placeholder="Connect" /><b-input  type="text" placeholder="host:port"/> </li>
  <li><h3>Language</h3>        host:port status (Excludes addpeers and peer database)</li>
  <li><h3>Language</h3><b-input type="checkbox" placeholder="Infrastructure mode" /></li>
      <li><b-input type="text" placeholder="host:port"/>
          <b-input  type="text" placeholder="username"/> <b-input  type="text" placeholder="password"/></li>
  <li><h3>TLS certificates</h3><b-input  type="textarea" placeholder="TLS public key"/></li>
  <li><h3>Language</h3><b-input  type="textarea" placeholder="TLS private key"/></li>
  <li><h3>Proxy</h3><b-input  type="text" placeholder="proxy url"/> </li>
  <li><h3>Language</h3><b-input  type="checkbox" placeholder="Tor" /></li>
  </ul>`
}
const Security = {
  template: `<ul class="flx fcl fii pdd stg">
  <li><h3>Change secret password</h3><b-input  type="text" placeholder="old password"/></li>
  <li><h3>new password</h3><b-input  type="text" placeholder="new password"/></li>
  <li><h3>again</h3><b-input  type="text" placeholder="again"/></li>
  <li><button class="btn">Change Public data encryption password</button></li>
  <li><h3>password</h3><b-input  type="text" placeholder="password"/></li>
  <li><h3>again</h3><b-input type="text" placeholder="again"/></li>
  </ul>`
}
const Mining = {
  template: `<ul class="flx fcl fii pdd stg">
  <li><h3>Language</h3><b-select placeholder="Algorithm"><option>SHA256D</option><option>Scrypt</option></b-select></li>
  <li>(Also sets what type of blocks expected from external miner)</li>
  <li><h3>CPU mine</h3><b-input type="text" placeholder="CPU mine" /></li>
  <li><h3>Cores</h3><b-input type="number" min="0" placeholder="xx cores" /></li>
  <li><button class="btn">Check for updates</button> <button class="btn">Install</button></li>
  </ul>`
}

const Settsssings = {
  template: `<div id="settings" class="flx fcl fii"><Title title="Settings" />
  <div class="flx fii">
  <b-form @submit.prevent="confirmSend" class="flx fcl fii panel">
  
  <b-tabs v-model="activeTab">
<b-tab-item label="Interface settings" class="flx fii pnlbg">
           <Interface />
</b-tab-item><b-tab-item label="Network settings" class="flx fii pnlbg">
            <Network />                  
</b-tab-item><b-tab-item label="Security settings" class="flx fii pnlbg">
            <Security />
</b-tab-item><b-tab-item label="Mining" class="flx fii pnlbg">
<Mining />
</b-tab-item></b-tabs>
<button type="submit" class="button is-primary" :disabled="!isFormValid">Send</button>
</b-form></div></div>`
}


const Settings = {
  template: `<div id="settings" class="flx fcl fii"><Title title="Settings" />
  <div class="flx fii">
  <b-tabs v-model="activeTab">
      <b-tab-item label="Interface settings" class="flx fii pnlbg"><Interface /></b-tab-item>
      <b-tab-item label="Network settings" class="flx fii pnlbg">Network </b-tab-item>
      <b-tab-item label="Security settings" class="flx fii pnlbg">Security </b-tab-item>
      <b-tab-item label="Mining" class="flx fii pnlbg">Mining </b-tab-item>
  </b-tabs>
  </div></div>`,
  data() {
      return {
          activeTab: 0,
      }
  }
}