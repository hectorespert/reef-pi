import React from 'react'
import JournalForm from './form'
import PropTypes from 'prop-types'
import { fetchJournalUsage } from '../redux/actions/journal'
import { connect } from 'react-redux'
import RecordJournalForm from './record_form'

class Journal extends React.Component {
  constructor (props) {
    super(props)
    this.handleOnSubmit = this.handleOnSubmit.bind(this)
  }

  componentDidMount () {
    this.props.fetchJournalUsage(this.props.data.id)
  }

  handleOnSubmit (values) {
    this.props.onSubmit(values)
  }

  render () {
    if (this.props.readOnly) {
      return (
        <div className='container'>
          <div className='col col-sm-6 col-md-3'>
            <RecordJournalForm />
          </div>
          <div className='col col-sm-6 col-md-3'>
            <p>Records</p>
          </div>
        </div>
      )
    } else {
      return (
        <JournalForm
          data={this.props.data}
          onSubmit={this.handleOnSubmit}
        />
      )
    }
  }
}

Journal.propTypes = {
  readOnly: PropTypes.bool,
  onSubmit: PropTypes.func.isRequired
}

const mapStateToProps = state => {
  return {
    journal_usage: state.journal_usage
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchJournalUsage: (id) => dispatch(fetchJournalUsage(id))
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Journal)
