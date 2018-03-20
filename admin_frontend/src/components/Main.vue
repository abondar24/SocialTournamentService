<template>
	<div class="main">
		<div>
			<b-modal id="modal1" title="Something went wrong" v-model="errorModal">
				<pre class="my-4">{{msg}}</pre>
				<p class="my-4">{{statusCode}}</p>
			</b-modal>
		</div>
		<b-content>
			<b-btn @click="status">test get status</b-btn>
		</b-content>
		<!-- Make sure to wrap b-tab in b-tabs -->
		<b-tabs> 

			<b-tab title="Players" active>
					<br>Players data
			</b-tab>
			<b-tab title="Tournaments">
					<br>Tournament data
			</b-tab>

		</b-tabs>
	</div>
</template>

<script>
export default {
  name: 'Main',
	/** [Vue warn]: Avoid mutating a prop directly since the value will be overwritten 
	whenever the parent component re-renders.
	Instead, use a data or computed property based on the prop's value. */
  // props: { 
  //   msg: String,
  //   statusCode: Number,
  //   serverStatusOk: {
  //       default: true,
  //       type: Boolean
  //   },
  // },
	data() {
		return {
			errorModal: false,
			msg: '',
			statusCode: null,
			serverStatusOk: false
		}
	},
	methods: {
		status() {
			this.$http.get(this.$hostname+'/status')
				.then(response => {
					console.log(response);
					this.msg = response.data.msg;
					this.statusCode = response.data.code;
					this.errorModal = false;
					if (this.statusCode!==200){
						this.serverStatusOk = false;
						this.errorModal = true;
					}

				}, error => {
					console.error('sersver error', error)
					this.msg = error.statusText || error.data;
					this.statusCode = error.statusCode || error.status;
					this.serverStatusOk = false;
					this.errorModal = true;
				});
		},
	},
	created: function () {
			this.status();
			console.log('THIS_HOSTNAME', this.$hostname)
	}
}
</script>

<style src="../styles/main.css">
</style>
