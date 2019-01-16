var AddressBookC = {
  template: vuedata.data.pages.addressbook,
  props:{
    vicons:Object,
    vlng:Object,
  },
  data () {
    return {
      address:"",
      label:"",
      addressbook: addressbook.data,
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
    labelForm: function() {
      addressbooklabel.addressBookLabelWrite(this.label, this.address);
     },
  confirmCustomDelete() {
      this.$dialog.confirm({
          title: 'Deleting label',
          message: 'Are you sure you want to <b>delete</b> your label? This action cannot be undone.',
          confirmText: 'Delete Label',
          type: 'is-danger',
          hasIcon: true,
          onConfirm: () => this.$toast.open('Label deleted!')
      })
  }
}
}