var MyVue = Vue.extend({
  delimiters: ['((', '))']
})

var fragmentsList = new MyVue({
  el: '#fragments-list',
  data: {
	fragments: []
  },
  methods: {
	update: function(e) {
	  axios.get('/api/fragments')
		.then(function(response){
		  fragmentsList.fragments = response.data
		})
		.catch(function(error){
		  console.log(error)
		})
	}
  }
})

fragmentsList.update()

