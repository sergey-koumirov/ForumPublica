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

function post(vm, url, payload, ignore, callback){
  vm.isLoading = true;
  axios.post(
      url,
      payload
  ).then(function(response){
      if(!ignore){
          vm.data = response.data;
          console.log(response.data);
      }
      if(callback){
          callback();
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
      chars: [],
      data: {},
      usedLocations:[],
      usedStores:[],
      showAddModal: false,
      showWhereModal: false,
      showStoreModal: false,
      isLoading: false,
      addModal: {
        selectedTypeId: null,
      },
      whereModal: {
        marketItemId: null,
        characterId: "",
        selectedLocationId: null,
        selectedLocationType: null,
      },
      storeModal: {
        marketItemId: null,
        characterId: "",
        selectedLocationId: null,
        selectedLocationType: null,        
      }
    },

    created: function () {
      // new Clipboard('.copy-item');
    },

    beforeMount: function () {
        this.chars = JSON.parse(this.$el.attributes['chars'].value);
        this.data = JSON.parse(this.$el.attributes['market-items'].value);
        this.data.Records.forEach((record)=>{
            record.Locations.forEach((location)=>{

                if( !this.usedLocations.find((el)=>{return el.LocationID == location.LocationID}) ){
                    this.usedLocations.push(location);
                }

                if( (location.Type=='station' || location.Type=='structure') && !this.usedStores.find((el)=>{return el.LocationID == location.LocationID}) ){
                    this.usedStores.push(location);
                }

            });
        });
    },

    methods: {
      OpenAddModal: function(){
        this.addModal.selectedTypeId = null;
        this.showAddModal = true;
      },

      OpenWhereModal: function(marketItemId){
        this.whereModal.marketItemId = marketItemId;        
        this.whereModal.selectedLocationId = null;
        this.showWhereModal = true;
      },

      OpenStoreModal: function(marketItemId){
        this.storeModal.marketItemId = marketItemId;        
        this.storeModal.selectedLocationId = null;
        this.showStoreModal = true;
      },

      CloseAddModal: function(){
        this.showAddModal = false;
      },

      CloseWhereModal: function(){
        this.showWhereModal = false;
      },

      CloseStoreModal: function(){
        this.showStoreModal = false;
      },

      TypeSelected: function(typeID){
          var self = this;
          self.addModal.selectedTypeId = typeID;
          if(!!typeID){
              post(this, '/app/market_items?page='+this.data.Page, {TypeID: typeID}, false, ()=>{ self.$emit('redraw') });
              this.showAddModal = false;
          }
      },

      LocationSelected: function(id,text,type){
        this.whereModal.selectedLocationId = id;
        this.whereModal.selectedLocationType = type;
      },

      onUsedLocationChange: function(event){
          var l = this.usedLocations.find((el)=>{return el.LocationID == event.target.value});
          if(!!l){
              this.whereModal.selectedLocationId = l.LocationID;
              this.whereModal.selectedLocationType = l.Type;
          }else{
              this.LocationReset();
          }
      },

      onUsedStoreChange: function(event){
          var l = this.usedStores.find((el)=>{return el.LocationID == event.target.value});
          if(!!l){
              this.storeModal.selectedLocationId = l.LocationID;
              this.storeModal.selectedLocationType = l.Type;
          }else{
              this.StoreReset();
          }
      },

      LocationReset: function(){
        this.whereModal.selectedLocationId = null;
      },

      StoreSelected: function(id,text,type){
        this.storeModal.selectedLocationId = id;
        this.storeModal.selectedLocationType = type;
      },
      StoreReset: function(){
        this.storeModal.selectedLocationId = null;
      },

      AddWhere: function(){
        post(
          this, 
          '/app/market_item/'+this.whereModal.marketItemId+'/locations', 
          {
            LocationID: this.whereModal.selectedLocationId,
            LocationType: this.whereModal.selectedLocationType,
            CharacterID: this.whereModal.characterId,
          }
        );
        this.CloseWhereModal();
      },

      AddStore: function(){
        post(
          this, 
          '/app/market_item/'+this.storeModal.marketItemId+'/stores', 
          {
            LocationID: this.storeModal.selectedLocationId,
            LocationType: this.storeModal.selectedLocationType,
            CharacterID: this.storeModal.characterId,
          }
        );
        this.CloseStoreModal();
      },


      DeleteMarketLocation: function(miId, lId){
        del(this, '/app/market_item/'+miId+'/location/'+lId+'?page='+this.data.Page);
      },
      
      DeleteMarketStore: function(miId, sId){
        del(this, '/app/market_item/'+miId+'/store/'+sId+'?page='+this.data.Page);
      },

      OpenCharPopup: function($event, typeId){
        this.$root.$emit('open-market', typeId, $event.pageX, $event.pageY)
        this.HighlightRow($event);
      },
    },
});


