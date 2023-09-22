<template>
  <div >
    <Search @updateTerm="term = $event" @searched="updateEmailContents" />
    <EmailList :term="term" :emailContents="emailContents" />
    </div>
</template>

<script>
import Search from './components/Search.vue';
import EmailList from './components/EmailList.vue'

export default {
  components: {
    Search,
    EmailList
  },
  data() {
    return {
      emailContents: [],
      term: ''
    };
  },
  methods: {
    updateEmailContents(newEmailContents) {
      this.emailContents = newEmailContents;
    }
  }
};
</script>
