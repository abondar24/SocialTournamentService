<template src="../templates/main.html"></template>

<script>
    export default {
  name: 'Main',
  props: {
    msg: String,
    statusCode: Number,
    serverStatusOk: {
        default: true,
        type: Boolean
    },
  },

    methods: {
        status: function () {
            this.$http.get(this.$hostname+'/status')
                .then(response => {
                    console.log(response);
                this.msg = response.data.msg;
                this.statusCode = response.data.code;
                    if (this.statusCode!==200){
                        this.serverStatusOk = false;
                    }
        },error=>{
                    this.msg = error.data;
                    this.statusCode = error.statusCode;
                    this.serverStatusOk = false;
                });
        },

    },
    created: function () {
        this.status();
    }
}
</script>

<style src="../styles/main.css">
</style>
