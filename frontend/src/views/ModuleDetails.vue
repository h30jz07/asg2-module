<template>
  <div>
    <h1>({{ code }}) Module Details</h1>
  <div v-if="link" class="text">
      <a :href=link>View Module Ratings & Comments</a>
  </div>
  <div v-else>
    <p>Loading ratings and comments link...</p>
  </div>
  <div id="tabNav">
    <TabNav>
      <Tab title="Students">
        <div v-if="students">
          <div>
            <h3>Enrolled Students</h3>
            <div v-for="student in students" :key="student.student_id" class="student">
              <p>Student ID: {{ student.student_id }}</p>
              <p>Class ID: {{ student.class_id }}</p>
            </div>
          </div>
        </div>
        <div v-else>
          <p>Loading module details...</p>
        </div>
      </Tab>

      <Tab title="Tutors">
        <div v-if="tutors">
          <div>
            <h3>Assigned Tutors</h3>
            <div v-for="tutor in tutors" :key="tutor.tutorid" class="student">
              <p>Tutor ID: {{ tutor.tutor_id }}</p>
              <p>Name: {{ tutor.name }}</p>
              <p>Email: {{ tutor.email }}</p>
              <p>Description: {{ tutor.descriptions }}</p>
            </div>
          </div>
        </div>
        <div v-else>
          <p>Loading module details...</p>
        </div>
      </Tab>

      <Tab title="Classes">
        <div v-if="classes"> 
          <div>
          <h3>Classes</h3>
          <div v-for="classItem in classes" :key="classItem.class_id" class="class">
            <p class="classItem">{{ classItem }}</p>
          </div>
          </div>
        </div>
        <div v-else>
          <p>Loading module details...</p>
        </div>
      </Tab>
    </TabNav>
  </div>
  </div>
</template>

<script>
import TabNav from '../components/TabNav.vue'
import Tab from '../components/Tab.vue'

export default {
    props: ['code'],
    components: { TabNav, Tab },
    methods: {
      setSelected(tab) {
        this.setSelected = tab;
      }
    },    
    data() {
        return{
          selected: 'Students',
          students: null,
          tutors: null,
          classes: null,
          link: null
        }
    },
    mounted() {
        fetch(`http://${process.env.VUE_APP_URL}:${process.env.VUE_APP_BACKEND_PORT}/module/v1/details/${this.code}`)
        //This section need to refactor to the new return from 3.7's golang backend
            .then(resp => resp.json())
            .then(data => {
              this.students = data.enrolled_students
              this.tutors = data.assigned_tutors
              this.classes = data.classes
              this.link = data.ratings_and_comments_url
            })
            .catch(err => console.log(err.message))
            console.log(this.classes)
    },
}
</script>

<style>
    p {
      margin: 0;
    }
  #tabNav{
    display: flex;
    flex-direction: row;
    justify-content: center;
  }
  .student{
      border: 2px solid#ddd;
      padding: 20px;
      border-radius: 5px;
      margin: 10px auto;
      max-width: 600px;
      color: #444;
      text-align: left;
    }
  .class{
      margin:0;
      border: 1px solid#ddd;
      padding-left: 20px;
      max-width: 600px;
      color: #444;
      text-align: left;
      position: relative
    }
    .classItem{
      top: 50%;
      margin: 2px;
    }
    .text {
      margin: 20px;
    }
    h1 {
        padding: 20px;
    }
</style>
