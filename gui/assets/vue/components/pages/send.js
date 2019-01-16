var SendC = {
  template: vuedata.data.pages.send,
  props:{
    vicons:Object,
    vlng:Object,

  },
  
  data () {
    return {
      address:"",
      label:"",
      passphrase:"",
      amount:0,
      timeout:220,
      transactions: blockchaindata.data.listallsendtransactions,
      isPaginated: true,
      isPaginationSimple: false,
      defaultSortDirection: 'asc',
      currentPage: 1,
      perPage: 5,
    }
  },
  components: {
},
  methods: {
     cardModal() {
      this.$modal.open({
          parent: this,
          component: ModalForm,
          hasModalCard: true
      })
  }
    },

}



var ModalForm = {
props: passphrase,address,label,amount,timeout,
methods: {
  sendForm() {
    // rpcinterface.walletPassphrase(this.passphrase, this.timeout);
    // rpcinterface.sendToAddress(this.address, this.label, this.amount);
   },
  },
template: `
    <form @submit.prevent="sendForm">
        <div class="modal-card" style="width: auto">
            <header class="modal-card-head">
            <p class="modal-card-title">Last verification!</p>
            <p>Send: <span v-model="amount"></span></p>
            <p>Label: <span v-model=""></span></p>
            <p>Label: <span v-model=""></span></p>
            <p>Label: <span :value="label"></span></p>
            <p>Label: <span v-model="timeout"></span></p>
            <p>To: <small v-model="address"></small></p>
            </header>
            <section class="modal-card-body">
          

                <b-field label="Passphrase">
                    <b-input
                        type="password"
                        password-reveal
                        :value="passphrase"
                        placeholder="Wallet Passphrase"
                        required>
                    </b-input>
                </b-field>
            </section>
            <footer class="modal-card-foot">
                <button class="button" type="button" @click="$parent.close()">Close</button>
                <button class="button is-primary">Send</button>
            </footer>
        </div>
    </form>
`
}
