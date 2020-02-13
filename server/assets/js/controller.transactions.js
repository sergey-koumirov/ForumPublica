var transactions = new Vue({
    el: '#transactions',

    data: {
      data: {},
      isLoading: false,
    },

    created: function () {
      // new Clipboard('.copy-item');
    },

    beforeMount: function () {
        this.data = JSON.parse(this.$el.attributes['transactions'].value);
    },

    methods: {
    },
});


