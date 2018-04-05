<template src="../templates/prize_form.html">
    
</template>

<script>

    import EventBus from './event-bus';
    export default {
        name: "UpdatePrizeForm",
        data(){
            return {
               winner: {
                   playerId: null,
                   prize:null
               }
            }
        },
        methods: {
            onSubmit(evt){
                evt.preventDefault();
                EventBus.$emit('prize',this.winner);
            }
        }
    }
</script>

<style scoped src="../styles/prize.css">

</style>