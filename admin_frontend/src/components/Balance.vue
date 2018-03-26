<template src="../templates/balance.html">
    
</template>

<script>
    import EventBus from './event-bus';
    export default {
        name: "Balance",
        data(){
            return {
                playerId: null,
                balanceErr:{
                    statusCode: 0,
                    errorMsg: '',
                    errorAlert: false
                },

                balance:''
            }

        },
        methods: {
            onSubmit(evt){
                evt.preventDefault();
                this.$http.get(this.$hostname+'/balance/'+this.playerId)
                    .then(response => {
                        this.balanceErr.statusCode = response.data.code;
                        if (this.balanceErr.statusCode!==200){
                            this.balanceErr.errorMsg = response.data;
                            this.balanceErr.errorAlert = true;
                            EventBus.$emit('balanceError',this.balanceErr);
                            this.playerId=null;
                        } else {
                            this.balance = response.data.msg.balance;
                            this.playerId=null;
                        }
                    }, error => {
                        this.balanceErr.errorMsg = error.statusText || error.data;
                        this.balanceErr.statusCode = error.statusCode || error.status;
                        this.balanceErr.errorAlert = true;
                        EventBus.$emit('balanceError',this.balanceErr);
                        this.playerId=null;
                    });
            },

        },
    }


</script>

<style scoped src="../styles/balance.css">

</style>