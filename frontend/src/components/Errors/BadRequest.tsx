const BadRequest = () => {
  return (
    <div className=" d-flex justify-content-center container text-center mt-5">
      <div className="row">
        <div className="col-md-12">
          <h1 className="display-1">400</h1>
          <h2>Bad Request</h2>
          <p className="lead">
            Sorry, we were unable to complete your request. Please try again.
          </p>
        </div>
      </div>
    </div>
  );
};

export default BadRequest;
