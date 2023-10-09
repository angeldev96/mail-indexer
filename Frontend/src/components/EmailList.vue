

<template>
  <div>

    <div 
      class="flex flex-col break-words border border-solid border-slate-100/[0.05] bg-gray-900 py-4 font-light text-gray-500 cursor-pointer"
      v-for="emailContent in emailContents"
      :key="emailContent._id"
      @click="selectEmail(emailContent)"
    >
      <div class="flex flex-grow items-center px-6">
        <div class="flex basis-[8.33333%] items-center">
          <div class="mr-2"></div>
          <div class="inline-block cursor-pointer text-slate-200"></div>
        </div>
        <div class="basis-[16.6667%] pr-3.5">
          <a href="#" class="flex cursor-pointer items-center text-slate-100">
            <div class="mr-4 flex h-10 w-10 items-center justify-center rounded-full bg-red-500 font-semibold text-white">
              <span>SA</span>
            </div>
            <span class="text-gray-200">{{ extractField(emailContent.content, 'From') }}</span>
          </a>
        </div>
        <div class="order-4 flex basis-[16.6667%] items-center justify-end">
          <div class="text-[0.80rem] text-gray-200">{{ formatDate(extractField(emailContent.content, 'Date')) }}</div>
          <div class="ml-4 flex align-middle"><button class="flex-grow cursor-pointer items-start rounded-lg border border-solid text-slate-900"></button></div>
        </div>
        <div class="flex basis-7/12 items-center pr-3.5">
          <a href="#" class="cursor-pointer overflow-hidden text-ellipsis">
            <span v-html="highlightKeyword(extractContentPreview(emailContent.content), term)"></span>
          </a>
        </div>
      </div>
    </div>

    <EmailDetail v-if="selectedEmail" :email="selectedEmail" :term="term" />
  </div>
</template>


<script>
import EmailDetail from './EmailDetail.vue';

export default {

  components: {
    EmailDetail,
    
  },
  props: {
    emailContents: {
      type: Array,
      required: true
    },
    term: {
    type: String,
    default: ''
  },
  totalHits: {  
      type: Number,
      default: 0
    },
  },
  
  data() {
    return {
      selectedEmail: null
    };
  },
  methods: {
    extractField(content, field) {
      if (!content) return '';
      const regex = new RegExp(`${field}: (.+)\\r\\n`);
      const match = content.match(regex);
      return match ? match[1].trim() : '';
    },

    extractContentPreview(content) {
      const body = content.split('\r\n\r\n')[1];
      return body ? body.split(' ').slice(0, 5).join(' ') + '...' : '';
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      return `${date.getDate().toString().padStart(2, '0')}/${(date.getMonth() + 1).toString().padStart(2, '0')}/${date.getFullYear()}`;
    },
    selectEmail(email) {
      this.selectedEmail = email;
    },
    highlightKeyword(text, keyword) {
      if (!keyword) return text;
      const regex = new RegExp(`(${keyword})`, 'gi');
      return text.replace(regex, '<span class="highlight">$1</span>');
    }

  }
};
</script>

<style scoped>
.highlight {
  background-color: yellow;
  padding: 2px;
}
</style>



