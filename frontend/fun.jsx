const students = [
    { name: 'Alice', grade: 82 },
    { name: 'Bob', grade: 90 },
    { name: 'Charlie', grade: 87 },
    { name: 'Diana', grade: 78 }
  ];

  const newArray = students.map((student, idx, array) => {
    const highestGrade = student.grade[0]
    for (var i = 1; i <= array.length; i++) {
        if (student.grade[i] > highestGrade) {
            highestGrade = student.grade[i]
        }
    }
    student.normalizedGrade = student.grade / highestGrade *100
    const newObject = {name: student.name, grade: student.grade, normalizedGrade: student.normalizedGrade}
    return newObject
  })