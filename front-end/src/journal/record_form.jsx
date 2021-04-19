import React from 'react'
import PropTypes from 'prop-types'
import { withFormik } from 'formik'
import JournalSchema from './schema'

const RecordJournal = () => {
  const handleSubmit = event => {
    event.preventDefault()
    console.log('Pepe')
  }
  return (
    <form onSubmit={handleSubmit}>
      <button type='submit' className='btn btn-primary mb-2'>Submit</button>
    </form>
  )
}

RecordJournal.propTypes = {
  values: PropTypes.object.isRequired,
  errors: PropTypes.object,
  touched: PropTypes.object,
  handleBlur: PropTypes.func.isRequired,
  submitForm: PropTypes.func.isRequired,
  onDelete: PropTypes.func,
  handleChange: PropTypes.func
}

const RecordJournalForm = withFormik({
  displayName: 'RecordJournalForm',
  mapPropsToValues: props => {
    let data = props.data
    if (data === undefined) {
      data = {
        name: '',
        description: '',
        unit: ''
      }
    }
    const values = {
      id: data.id || '',
      name: data.name || '',
      description: data.description || '',
      unit: data.unit || ''
    }
    return values
  },
  validationSchema: JournalSchema,
  handleSubmit: (values, { props }) => {
    props.onSubmit(values)
  }
})(RecordJournal)

export default RecordJournalForm
