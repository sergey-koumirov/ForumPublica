function patch(vm, url, payload, ignore){
  vm.isLoading = true,
  axios.patch(
      url,
      payload
  ).then(function(response){
      if(!ignore){
          vm.construction = response.data;
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
          vm.construction = response.data;
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
          vm.construction = response.data;
      }
      vm.isLoading = false;
  }).catch(function(error){
      console.log(error);
  });
}


var constructions = new Vue({
    el: '#construction',

    data: {
      construction: {},
      chars: [],
      isLoading: false,
      showMEModal: false,
      showQtyModal: false,
      showRunModal: false,
      showExpenseModal: false,
      ask: {
        me: null,
        bpId: null,
        qty: null,
        exValue: null,
        description: null,
        repeats: null,
      }
    },

    created: function () {
      new Clipboard('.copy-item');
    },

    beforeMount: function () {
        this.construction = JSON.parse(this.$el.attributes['construction'].value);
        this.chars = JSON.parse(this.$el.attributes['chars'].value);
    },

    methods: {

        SaveBonus: function(){
            post(
                this,
                '/app/construction/'+this.construction.Model.ID+'/save_bonus',
                {
                    CitadelType: this.construction.Model.CitadelType,
                    RigFactor: this.construction.Model.RigFactor,
                    SpaceType: this.construction.Model.SpaceType,
                },
                false
            )
        },

        TypeSelected: function(typeId){
          post(this, '/app/construction/'+this.construction.Model.ID+'/add_blueprint', {BlueprintId: typeId})
        },

        DeleteBpo: function(cnBpoId){
          if (confirm("Sure?")) {
            del(this, '/app/construction/'+this.construction.Model.ID+'/blueprint/'+cnBpoId)
          }
        },

        ChangeME: function(bpId, me){
          this.ask.me = me;
          this.ask.bpId = bpId;
          this.showMEModal = true;
        },

        SaveME: function(){
          patch(
            this,
            '/app/construction/'+this.construction.Model.ID+'/blueprint/'+this.ask.bpId,
            {me: this.ask.me}
          )

          this.showMEModal = false;
        },

        ChangeQty: function(bpId, qty){
          this.ask.qty = qty;
          this.ask.bpId = bpId;
          this.showQtyModal = true;

          this.$nextTick(function () {
            $('.bpo-modal .qty').focus().select();
          }); 
        },
        

        SaveQty: function(){
          patch(
            this,
            '/app/construction/'+this.construction.Model.ID+'/blueprint/'+this.ask.bpId,
            {qty: this.ask.qty}
          );

          this.showQtyModal = false;
        },

        CloseQty: function(eventName){
          this.showQtyModal = false;
        },

        OpenCharPopup: function($event, typeId){
          this.$root.$emit('open-market', typeId, $event.pageX, $event.pageY)
          this.HighlightRow($event);
        },

        OpenRunModal: function(bpo){

          console.log(bpo.SgtRepeats, bpo.SgtRunQty);

          this.ask.bpId = bpo.Model.TypeID;
          this.ask.repeats = bpo.SgtRepeats;
          this.ask.qty = bpo.DefaultME == 2 ? 10 : bpo.SgtRunQty;
          this.ask.me = bpo.DefaultME;
          this.showRunModal = true;
          this.$nextTick(function () {
            $('.run-modal .repeats').focus().select();
          }); 
        },

        SaveRun: function(){
          post(
            this, 
            '/app/construction/'+this.construction.Model.ID+'/add_run', 
            {
              me: this.ask.me, 
              repeats: this.ask.repeats, 
              qty: this.ask.qty, 
              BlueprintId: this.ask.bpId
            }, 
            false
          );
          this.showRunModal = false;
        },

        CloseRun: function(eventName){
          this.showRunModal = false;
        },

        DeleteRun: function(runId){
            del(
                this,
                '/app/construction/'+this.construction.Model.ID+'/run/'+runId,
                false
            );
        },

        SetExcluded: function(item){
          item.Excluded = true;
        },

        ResetExcluded: function(item){
            item.Excluded = false;
        },

        HasExcluded: function(){
            if(!!this.construction.Materials){
                for(i = 0; i < this.construction.Materials.length; i++){
                    if(this.construction.Materials[i].Excluded){ return true; }
                }
            }
            return false;
        },

        HighlightRow: function($event){
            $($event.currentTarget).closest("table").find("tr").removeClass("highlighted");
            $($event.currentTarget).closest("tr").addClass("highlighted");
        },

        PartialVol: function(){
            var result = 0;
            this.construction.Materials.forEach(function(item){
                if(!item.Excluded){
                    result = result + item.Volume||0;
                }
            });
            return result;
        },

        FullVol: function(){
            var result = 0;
            this.construction.Materials.forEach(function(item){
                result = result + item.Volume||0;
            });
            return result;
        },

        PartialPrice: function(){
            var result = 0;
            this.construction.Materials.forEach(function(item){
                if(!item.Excluded){
                    result = result + item.Price||0;
                }
            });
            return result;
        },

        FullCost: function(){
            var result = 0;
            this.construction.Materials.forEach(function(item){
                result = result + (item.Price||0) * (item.Qty||0);
            });
            return result;
        },

        TotalExpenses: function(bpo){
            var result = 0;
            bpo.Expenses.forEach(function(item){
                result = result + item.ExValue||0;
            });
            return result;
        },

        AvgExpenses: function(bpo){
            var result = 0;
            bpo.Expenses.forEach(function(item){
                result = result + item.ExValue||0;
            });
            return result / bpo.Model.Qty;
        },

        OpenExpenseModal: function(bpo){
            this.ask.exValue = 0;
            this.ask.description = null;
            this.ask.bpId = bpo.Model.ID;
            this.showExpenseModal = true;
            this.$nextTick(function () {
                $('input.expense').focus().select();
            });         
        },

        CloseExpense: function(eventName){
            this.showExpenseModal = false;
        },

        SaveExpense: function(){
            post(
              this, 
              '/app/construction/'+this.construction.Model.ID+'/expenses', 
              {
                Description: this.ask.description, 
                ExValue: this.ask.exValue, 
                BpoId: this.ask.bpId,
              }, 
              false
            );
            this.showExpenseModal = false;
        },
        DeleteExpense: function(id){
            del(
              this, 
              '/app/construction/'+this.construction.Model.ID+'/expense/'+id, 
              false
            );
        },
    },
});


