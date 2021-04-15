<template>
  <data-tables
    :data='[]'
    :total='100'>
    <div slot="empty" style='color: red'>emptyyyyyyyyyyyyyy</div>
    <el-table-column v-for="title in titles"
      :prop="title.prop"
      :label="title.label"
      :key="title.label"
      sortable="custom">
    </el-table-column>
  </data-tables>
</template>


<script>
export default {
  name: 'tables',
  data() {
    return {
      data: {
            titles: [{
            prop: "hi", 
            label: "hello"
        }, 
        {
            prop: "hii", 
            label: "hellooo"
        }]
     }
      
    }
  }
}
</script>
