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
          vm.data = response.data;
          console.log(response.data);
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
        vm.data = response.data;
      }
      vm.isLoading = false;
  }).catch(function(error){
      console.log(error);
  });
}


var marketItems = new Vue({
    el: '#market-items',

    data: {
      data: {},
      showAddModal: false,
      showWhereModal: false,
      isLoading: false,
      addModal: {
        selectedTypeId: null,
      },
      whereModal: {
        marketItemId: null,
        characterId: "",
        selectedLocationId: null,
        selectedLocationType: null,
        selectedWarehouseId: null,
        selectedWarehouseType: null,
      }
    },

    created: function () {
      // new Clipboard('.copy-item');
    },

    beforeMount: function () {
        this.data = JSON.parse(this.$el.attributes['market-items'].value);
    },

    methods: {
      OpenAddModal: function(){
        this.addModal.selectedTypeId = null;
        this.showAddModal = true;
      },

      OpenWhereModal: function(marketItemId){
        this.whereModal.marketItemId = marketItemId;        
        this.whereModal.selectedLocationId = null;
        this.whereModal.selectedWarehouseId = null;
        this.showWhereModal = true;
      },

      CloseAddModal: function(){
        this.showAddModal = false;
      },

      CloseWhereModal: function(){
        this.showWhereModal = false;
      },

      TypeSelected: function(typeID){
        this.addModal.selectedTypeId = typeID;
        if(!!typeID){
          post(this, '/app/market_items?page='+this.data.Page, {TypeID: typeID});
          this.showAddModal = false;
        }        
      },

      LocationSelected: function(id,text,type){
        this.whereModal.selectedLocationId = id;
        this.whereModal.selectedLocationType = type;
      },
      LocationReset: function(){
        this.whereModal.selectedLocationId = null;
      },
      WarehouseSelected: function(id,text,type){
        this.whereModal.selectedWarehouseId = id;
        this.whereModal.selectedWarehouseType = type;
      },
      WarehouseReset: function(){
        this.whereModal.selectedWarehouseId = null;
      },

      AddWhere: function(){
        //todo insert
        post(
          this, 
          '/app/market_item/'+this.whereModal.marketItemId+'/locations', 
          {
            LocationID: this.whereModal.selectedLocationId,
            LocationType: this.whereModal.selectedLocationType,
            StoreLocationID: this.whereModal.selectedWarehouseId,
            StoreLocationType: this.whereModal.selectedWarehouseType,
            CharacterID: this.whereModal.characterId,
          }
        );
        this.CloseWhereModal();
      },

      DeleteMarketLocation: function(miId, lId){
        del(this, '/app/market_item/'+miId+'/location/'+lId+'?page='+this.data.Page);
      }

    },
});


