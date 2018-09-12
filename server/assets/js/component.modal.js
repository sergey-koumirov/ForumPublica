Vue.component('modal', {
    template: '  <div class="modal-mask">'+
              '    <div class="modal-wrapper">'+
              '      <div class="modal-container">'+
              '        <div class="modal-body">'+
              '          <slot name="body">default body</slot>'+
              '        </div>'+
              '        <div class="modal-footer">'+
              '          <slot name="footer">'+
              '            <button class="btn btn-primary btn-sm modal-default-button" @click="$emit(\'close\')">Close</button>'+
              '          </slot>'+
              '        </div>'+
              '      </div>'+
              '    </div>'+
              '  </div>',
})
