const InternalServerError = () => {
  return (
    <div className=" d-flex justify-content-center container text-center mt-5">
      <div className="row">
        <div className="col-md-12">
          <h1 className="display-1">500</h1>
          <h2>Internal Server Error</h2>
          <p className="lead">Sorry, something happened from our side.</p>
        </div>
      </div>
    </div>
  );
};

export default InternalServerError;
