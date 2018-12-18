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
                '/app/construction/'+this.construction.Model.Id+'/save_bonus',
                {
                    CitadelType: this.construction.Model.CitadelType,
                    RigFactor: this.construction.Model.RigFactor,
                    SpaceType: this.construction.Model.SpaceType,
                },
                true
            )
        },

        TypeSelected: function(typeId){
          post(this, '/app/construction/'+this.construction.Model.Id+'/add_blueprint', {BlueprintId: typeId})
        },

        DeleteBpo: function(cnBpoId){
          if (confirm("Sure?")) {
            del(this, '/app/construction/'+this.construction.Model.Id+'/blueprint/'+cnBpoId)
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
            '/app/construction/'+this.construction.Model.Id+'/blueprint/'+this.ask.bpId,
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
            '/app/construction/'+this.construction.Model.Id+'/blueprint/'+this.ask.bpId,
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
          this.ask.bpId = bpo.Model.TypeId;
          this.ask.repeats = 1;
          this.ask.qty = bpo.DefaultME == 2 ? 10 : 1;
          this.ask.me = bpo.DefaultME;
          this.showRunModal = true;
          this.$nextTick(function () {
            $('.run-modal .repeats').focus().select();
          }); 
        },

        SaveRun: function(){
          post(
            this, 
            '/app/construction/'+this.construction.Model.Id+'/add_run', 
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
                '/app/construction/'+this.construction.Model.Id+'/run/'+runId,
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

        FullPrice: function(){
            var result = 0;
            this.construction.Materials.forEach(function(item){
                result = result + item.Price||0;
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
            this.ask.bpId = bpo.Model.Id;
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
              '/app/construction/'+this.construction.Model.Id+'/expenses', 
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
              '/app/construction/'+this.construction.Model.Id+'/expense/'+id, 
              false
            );
        },
    },
});



// var constructionApp = angular.module('constructionApp', ['angular-jquery-autocomplete','ngclipboard']);

// constructionApp.controller('ConstructionCtrl', function ($scope, $http, $timeout, $location) {
//
//     $scope.newId = null;
//     $scope.newText = null;
//     $scope.isLoading = false;
//     $scope.constructionId = null;
//     $scope.construction = {};
//     $scope.constructionStores = [];
//     $scope.storeName = "";
//     $scope.storeId = null;
//     $scope.pasteBinText = "";
//     $scope.MarketTypeId = null;
//     $scope.ask = {
//         Number: 0,
//         Description: "",
//         Func: null,
//         TypeId: null,
//         CnBpoId: null,
//         Error: null,
//         ME: 10,
//         TE: 20,
//         Status: "planned",
//         Repeat: 1,
//         Comment: '',
//     };
//
//     $scope.init = function(constructionId){
//         $scope.constructionId = constructionId;
//         $scope.isLoading = true;
//         $http.get(
//             '/construction/'+$scope.constructionId+'.json'
//         ).success(function(data) {
//             $scope.construction = data;
//             $scope.isLoading = false;
//
//             if(!$scope.construction.Store){
//                 $scope.isLoading = true;
//                 $http.get(
//                     '/construction/stores.json'
//                 ).success(function(data){
//                     $scope.constructionStores = data;
//                     $scope.isLoading = false;
//                 });
//             }
//
//         });
//     };
//
//
//     $scope.$watch('newId', function(newValues, oldValues, scope) {
//         if(!!$scope.newId){
//             $scope.AddBpo($scope.newId);
//             $scope.newId = null;
//             $scope.newText = null;
//         }
//     });
//
//     $scope.AddBpo = function(bpoId){
//         $scope.isLoading = true;
//         $http.get(
//             '/construction/'+$scope.constructionId+'/add_bpo.json?bpo_id='+bpoId
//         ).success(function(data) {
//             $scope.construction = data;
//             $scope.isLoading = false;
//         });
//     };
//
//
//     $scope.CreateStore = function(){
//         if (!!$scope.storeName && confirm("Sure? Store can't be unlinked.")) {
//             $scope.isLoading = true;
//             $http.get(
//                 '/construction/'+$scope.constructionId+'/new_store.json?name='+$scope.storeName
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
//
//     $scope.LinkStore = function(){
//         if (!!$scope.storeId && confirm("Sure? Store can't be unlinked.")) {
//             $scope.isLoading = true;
//             $http.get(
//                 '/construction/'+$scope.constructionId+'/link_store.json?store_id='+$scope.storeId
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
//
//     $scope.ReplaceStoreList = function(){
//         $scope.ask.pasteBinText = "";
//         $scope.ask.Func = function(){
//             $scope.isLoading = true;
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/replace_store_items.json',
//                 {data: $scope.pasteBinText}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//                 $('#storeModal').modal('hide');
//             });
//         };
//     };
//
//     $scope.AddStoreList = function(){
//         $scope.ask.pasteBinText = "";
//         $scope.ask.Func = function(){
//             $scope.isLoading = true;
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/add_store_items.json',
//                 {data: $scope.pasteBinText}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//                 $('#storeModal').modal('hide');
//             });
//         };
//     };
//
//     $scope.AddStoreItemQtyAskNumber = function(typeId){
//         $scope.ask.Number = 0;
//         $scope.ask.TypeId = typeId;
//         $scope.ask.Func = function(){
//             $scope.PostUrlWithLoading('/construction/'+$scope.constructionId+'/add_store_item_qty.json', $scope.ask.Number, $scope.ask.TypeId);
//         };
//     };
//
//     $scope.ChangeStoreItemQtyAskNumber = function(typeId, current){
//         $scope.ask.Number = current;
//         $scope.ask.TypeId = typeId;
//         $scope.ask.Func = function(){
//             $scope.PostUrlWithLoading('/construction/'+$scope.constructionId+'/change_store_item_qty.json', $scope.ask.Number, $scope.ask.TypeId);
//         };
//     };
//
//     $scope.ChangeBpoQtyAskNumber = function(constructionBpoId, current){
//         $scope.ask.Number = current;
//         $scope.ask.CnBpoId = constructionBpoId;
//         $scope.ask.Func = function(){
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/change_bpo_qty.json',
//                 {cn_bpo_id: $scope.ask.CnBpoId, qty: parseInt($scope.ask.Number)}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//                 $('#askNumberModal').modal('hide');
//             });
//         };
//     };
//
//     $scope.AddExpense = function(){
//         $scope.ask.Number = 0;
//         $scope.ask.Description = "";
//         $scope.ask.Func = function(){
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/add_expense.json',
//                 {
//                     value: $scope.ask.Number,
//                     description: $scope.ask.Description,
//                     cn_bpo_id: parseInt($scope.ask.CnBpoId)
//                 }
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//                 $('#askExpenseModal').modal('hide');
//             });
//         };
//     };
//
//     $scope.DeleteExpense = function(id){
//         if (confirm("Sure?")) {
//             $scope.isLoading = true;
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/delete_expense.json',
//                 {id: id}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
//
//     $scope.AddBpoQtyAskNumber = function(constructionBpoId){
//         $scope.ask.Number = 1;
//         $scope.ask.CnBpoId = constructionBpoId;
//         $scope.ask.Func = function(){
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/add_bpo_qty.json',
//                 {cn_bpo_id: $scope.ask.CnBpoId, qty: parseInt($scope.ask.Number)}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//                 $('#askNumberModal').modal('hide');
//             });
//         };
//     };
//
//     $scope.DeleteStoreItem = function(typeId){
//         if (confirm("Sure?")) {
//             $scope.isLoading = true;
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/delete_store_item.json',
//                 {type_id: typeId}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
//
//     $scope.ClearStore = function(){
//         if (confirm("Sure?")) {
//             $scope.isLoading = true;
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/clear_store.json'
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
//
//     $scope.StartMnf = function(bpoTypeId, ME, TE){
//         $scope.ask.Number = 1;
//         $scope.ask.Repeat = 1;
//         $scope.ask.TypeId = bpoTypeId;
//         $scope.ask.ME = ME;
//         $scope.ask.TE = TE;
//         $scope.ask.Cost = 0;
//         $scope.ask.Error = null;
//
//         $scope.ask.CitadelType = $scope.construction.CitadelType;
//         $scope.ask.RigFactor = $scope.construction.RigFactor;
//         $scope.ask.SpaceType = $scope.construction.SpaceType;
//
//         $scope.ask.Comment = $scope.construction.Comment;
//
//         $scope.ask.Func = function(){
//             $scope.PostBpoRunsUrlWithLoading($scope.ask);
//         };
//     }
//
//     $scope.StartRunMnf = function(cnBpoRunId){
//         $scope.isLoading = true;
//         $http.post(
//             '/construction/'+$scope.constructionId+'/start_run.json',
//             {run_id: cnBpoRunId}
//         ).success(function(data){
//             $scope.construction = data;
//             $scope.isLoading = false;
//         });
//     }
//
//     $scope.ChangeME = function(constructionBpoId, me){
//         $scope.ask.ME = me;
//         $scope.ask.CnBpoId = constructionBpoId;
//         $scope.isLoading = true;
//         $scope.ask.Func = function(){
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/change_bpo_me.json',
//                 {cn_bpo_id: $scope.ask.CnBpoId, me: parseInt($scope.ask.ME)}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//                 $('#askMEModal').modal('hide');
//             });
//         };
//     }
//
//     $scope.DeleteRun = function(runId){
//         $scope.isLoading = true;
//         $http.post(
//             '/construction/'+$scope.constructionId+'/delete_bpo_run.json',
//             {run_id: runId}
//         ).success(function(data){
//             $scope.construction = data;
//             $scope.isLoading = false;
//         });
//     };
//
//     $scope.DeleteRuns = function(runs){
//         if (confirm("Sure?")) {
//             $scope.isLoading = true;
//             $http.post(
//                 '/construction/'+$scope.constructionId+'/delete_bpo_runs.json',
//                 {run_ids: runs.map(function(r){return r.ConstructionBpoRunId;})}
//             ).success(function(data){
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
//
//     $scope.FinishMnf = function(runId){
//         $scope.isLoading = true;
//         $http.post(
//             '/construction/'+$scope.constructionId+'/finish_bpo_run.json',
//             {run_id: runId}
//         ).success(function(data){
//             $scope.construction = data;
//             $scope.isLoading = false;
//         });
//     };
//
//     $scope.HighlightRow = function($event){
//         $($event.currentTarget).closest("table").find("tr").removeClass("highlighted");
//         $($event.currentTarget).closest("tr").addClass("highlighted");
//     };
//

//
//     $scope.MaterialsText = function(){
//         if (!$scope.construction.Materials){
//             return "";
//         }
//         var result = "";
//         angular.forEach($scope.construction.Materials.Items, function(item){
//             result = result + item.Type.name + ", " + item.Qty + "\n";
//         });
//         return result;
//     };
//
//     $scope.AddPhantomTrans = function(bpoId){
//         $scope.isLoading = true;
//         $http.post(
//             '/construction/'+$scope.constructionId+'/add_phantom_trans.json',
//             {bpo_id: bpoId}
//         ).success(function(data){
//             $scope.construction = data;
//             $scope.isLoading = false;
//         });
//     };
//
//     $scope.DeletePhantomTrans = function(bpoId){
//         $scope.isLoading = true;
//         $http.post(
//            '/construction/'+$scope.constructionId+'/delete_phantom_trans.json',
//             {bpo_id: bpoId}
//         ).success(function(data){
//             $scope.construction = data;
//             $scope.isLoading = false;
//         });
//     };
//
//     $scope.OpenCharPopup = function($event, typeId){
//         $scope.MarketTypeId = typeId;
//         $scope.HighlightRow($event);
//         $(".popup").show().offset({left: $event.pageX, top: $event.pageY});
//     };
//
//     $scope.OpenMarket = function(charId){
//         $http.post(
//             '/ui/open_market.json',
//             {type_id: $scope.MarketTypeId, cid: charId}
//         ).then(
//             function(response){
//                 $(".popup").hide();
//             },
//             function(response){
//                 $(".popup").hide();
//                 console.log(response);
//             }
//         );
//     };
//
//     //--private---------------------------------
//
//     $scope.PostUrlWithLoading = function(url, qty, typeId){
//         $scope.isLoading = true;
//         $http.post(
//             url, {qty: qty, type_id: typeId}
//         ).success(function(data){
//             $scope.construction = data;
//             $scope.isLoading = false;
//             $('#askNumberModal').modal('hide');
//         });
//     }
//
//     $scope.PostBpoRunsUrlWithLoading = function(ask){
//         $scope.isLoading = true;
//         $http.post(
//             '/construction/'+$scope.constructionId+'/add_bpo_run.json',
//             {
//                 cnt: ask.Number,
//                 repeat: ask.Repeat,
//                 bpo_type_id: ask.TypeId,
//                 me: ask.ME,
//                 te: ask.TE,
//                 cost: ask.Cost,
//                 citadel_type: ask.CitadelType,
//                 rig_factor: ask.RigFactor,
//                 space_type: ask.SpaceType,
//                 status: ask.Status,
//                 comment: ask.Comment
//              }
//         ).then(
//             function(response){
//                 $scope.construction = response.data;
//                 $scope.isLoading = false;
//                 $('#askBpoRunsModal').modal('hide');
//             },
//             function(response){
//                 $scope.isLoading = false;
//                 $scope.ask.Error = response.data.error;
//             }
//         );
//     };
//
//     $scope.FormatSeconds = function(seconds){
//         var ss = seconds % 60,
//             mm = Math.trunc(seconds/60) % 60,
//             hh = Math.trunc(seconds/(60*60)) % 24,
//             dd = Math.trunc(seconds/(60*60*24));
//
//         var result = ("0" + hh).slice(-2) + ":" + ("0" + mm).slice(-2) + ":" +("0" + ss).slice(-2);
//         if( dd > 0){ result = dd + "d " + result; }
//         return result;
//     };
//
//     $scope.SetExcluded = function(item){
//         item.excluded = true;
//     };
//
//     $scope.ResetExcluded = function(item){
//         item.excluded = false;
//     };
//
//     $scope.HasExcluded = function(){
//         if(!!$scope.construction.Materials && !!$scope.construction.Materials.Items){
//             for(i = 0; i < $scope.construction.Materials.Items.length; i++){
//                 if($scope.construction.Materials.Items[i].excluded){ return true; }
//             }
//         }
//         return false;
//     };
//
//     $scope.PartialVol = function(){
//         var result = 0;
//         angular.forEach($scope.construction.Materials.Items, function(item){
//             if(!item.excluded){
//                 result = result + item.Volume;
//             }
//         });
//         return result;
//     };
//
//     $scope.PartialPrice = function(){
//         var result = 0;
//         angular.forEach($scope.construction.Materials.Items, function(item){
//             if(!item.excluded){
//                 result = result + item.Price;
//             }
//         });
//         return result;
//     };
//
// });
//
// $(function(){
//     $(".popup .close").click(function(e){
//         $(this).closest(".popup").hide();
//     });
// });
