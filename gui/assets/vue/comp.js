const Title = {
  template:`<div id="title"><p>{{title}}</p></div>`,
  props: ['title'],
};

const HomeC = {
  template: `<div class="flx fcl fii"><div class="flx frw fii pdb"><div class="flx fcl fii mrl pnlbg"><div class="flx fii pdd wst"><ul class="stg">
      <li class="htg"><div v-html="vicons.balance"></div><h2>Balance: </h2><strong> 556.5656 DUO</strong></li>
      <li class="htg"><div v-html="vicons.unconfirmed"></div><h2>Unconfirmed: </h2><strong> 26.656 DUO</strong></li>
      <li class="htg"><div v-html="vicons.txnumber"></div><h2>Number of transaction: </h2> <strong>5</strong></li>
</ul></div></div><div class="flx fcl fii lts">
       <h2 class="pds">Latest Transactions</h2>
      <div class="flx fii" v-bar><ul class="mrr">
<li class="htg"><div class="flx fii frw pnlbg move"><div class="flx left-side"><div v-html="vicons.received"></div><div v-html="vicons.sent"></div></div><div class="flx frw fii right-side"><div class="date"><p>datum 21.212.22</p></div><div class="bottom">walletAlkajhsdklhjalkjhsd</div><div class="amount">15.215<span class="bolder"> DUO</span></div></div></div></li>
</ul></div></div></div></div>`,
  props:{
    vicons:Object,
  }
}

const SendC = {
  template: `<div>tessssrr  </div>`
}

const ReceiveC = {
  template: `<div>terrrrrrrrrrr  </div>`
}

const AddressBookC = {
  template: `<div>terr  </div>`
}

const HistoryC = {
  template: `<div>terr  </div>`
}