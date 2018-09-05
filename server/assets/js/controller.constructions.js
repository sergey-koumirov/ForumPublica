var constructions = new Vue({
    el: '#construction',

    data: {
      construction: {},
      isLoading: false,
    },

    beforeMount: function () {
        this.construction = JSON.parse(this.$el.attributes['construction'].value);
        console.log(this.construction);
    },

    methods: {

        SaveBonus: function(){
            var vm=this;
            vm.isLoading = true,
            axios.post(
                '/app/construction/'+this.construction.Id+'/save_bonus',
                {
                    CitadelType: this.construction.CitadelType,
                    RigFactor: this.construction.RigFactor,
                    SpaceType: this.construction.SpaceType,
                }
            ).then(function(response){
                vm.isLoading = false;
            }).catch(function(error){
                console.log(error);
            });

        },

    },

    filters: {
        seconds: function(seconds){
            var s = seconds||0,
                ss = s % 60,
                mm = Math.trunc(s/60) % 60,
                hh = Math.trunc(s/(60*60)) % 24,
                dd = Math.trunc(s/(60*60*24));

            var result = ("0" + hh).slice(-2) + ":" + ("0" + mm).slice(-2) + ":" +("0" + ss).slice(-2);
            if( dd > 0){ result = dd + "d " + result; }
            return result;
        },
    }
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
//     $scope.DeleteBpo = function(cnBpoId){
//         if (confirm("Sure?")) {
//             $scope.isLoading = true;
//             $http.get(
//                 '/construction/'+$scope.constructionId+'/delete_bpo.json?cn_bpo_id='+cnBpoId
//             ).success(function(data) {
//                 $scope.construction = data;
//                 $scope.isLoading = false;
//             });
//         }
//     };
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