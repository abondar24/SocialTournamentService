<template src="../templates/main.html">

</template>

<script>
    import Player from './Player.vue'
    import Tournament from './Tournament.vue'
export default {
  name: 'Main',
    components:{
      Player,
        Tournament
    },
	/** [Vue warn]: Avoid mutating a prop directly since the value will be overwritten
	whenever the parent component re-renders.
	Instead, use a data or computed property based on the prop's value. */
	data() {
		return {
			errorModal: false,
			msg: '',
			statusCode: null,
		}
	},
	methods: {
		status() {
			this.$http.get(this.$hostname+'/status')
				.then(response => {
					this.msg = response.data.msg;
					this.statusCode = response.data.code;
					this.errorModal = this.statusCode !== 200;
				}, error => {
                    this.msg = error.statusText || error.data;
                    this.statusCode = error.statusCode || error.status;
                    this.errorModal = true;
				});
		},
	},
	created: function () {
			this.status();
	}
}
</script>

<style scoped src="../styles/main.css">
</style>
