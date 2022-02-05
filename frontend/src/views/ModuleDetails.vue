<template>
  <div v-if="students">
    <h1>Module Details Page</h1>
    <p>The module code is {{ code }}</p>
    <div>
      <h3>Enrolled Students</h3>
      <div v-for="student in students" :key="student.student_id" class="student">
        <p>Student ID: {{ student.student_id }}</p>
        <p>Class ID: {{ student.class_id }}</p>
      </div>
    </div>
    <div>
      <h3>Assigned Tutors</h3>
      <div v-for="tutor in tutors" :key="tutor.tutorid" class="student">
        <p>Tutor ID: {{ tutor.tutorid }}</p>
      </div>
    </div>
    <div>
      <h3>Classes</h3>
      <div v-for="classItem in classes" :key="classItem.class_id" class="student">
        <p>Class ID: {{ classItem }}</p>
      </div>
    </div>
    <div>
      <a :href=link>View Module Ratings & Comments</a>
    </div>
  </div>
  <div v-else>
    <p>Loading module details...</p>
  </div>
</template>

<script>
export default {
    props: ['code'],
    data() {
        return{
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
    }
}
</script>

<style>
  .student{
        background: #f4f4f4;
        padding: 20px;
        border-radius: 10px;
        margin: 10px auto;
        max-width: 600px;
        cursor: pointer;
        color: #444;
    }
    .student a{
        text-decoration: none;
    }
</style>
