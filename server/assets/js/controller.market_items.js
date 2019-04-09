function patch(vm, url, payload, ignore){
  vm.isLoading = true,
  axios.patch(
      url,
      payload
  ).then(function(response){
      if(!ignore){
          // vm.construction = response.data;
      }
      vm.isLoading = false;
  }).catch(function(error){
      console.log(error);
  });
}

function post(vm, url, payload, ignore){
  vm.isLoading = true,
  axios.post(
      url,
      payload
  ).then(function(response){
      if(!ignore){
          // vm.construction = response.data;
      }
      vm.isLoading = false;
  }).catch(function(error){
      console.log(error);
  });
}

function del(vm, url, ignore){
  vm.isLoading = true,
  axios.delete(
      url
  ).then(function(response){
      if(!ignore){
          // vm.construction = response.data;
      }
      vm.isLoading = false;
  }).catch(function(error){
      console.log(error);
  });
}


var marketItems = new Vue({
    el: '#market-items',

    data: {
      marketItems: {},
      showAddModal: false,
      isLoading: false,
    },

    created: function () {
      // new Clipboard('.copy-item');
    },

    beforeMount: function () {
        this.marketItems = JSON.parse(this.$el.attributes['market-items'].value);

        console.log(this.marketItems);
    },

    methods: {
      OpenAddModal: function(){
        this.showAddModal = true;
      },

      CloseAddModal: function(){
        this.showAddModal = false;
      },

      TypeSelected: function(typeID){
        console.log(typeID);
        post(this, '/app/market_items', {TypeID: typeID}, false)
      },

    },
});


