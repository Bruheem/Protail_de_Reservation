const Unauthorized = () => {
  return (
    <div className=" d-flex justify-content-center container text-center mt-5">
      <div className="row">
        <div className="col-md-12">
          <h1 className="display-1">401</h1>
          <h2>Unauthorized Access</h2>
          <p className="lead">
            Sorry, you do not have permission to access this page.
          </p>
        </div>
      </div>
    </div>
  );
};

export default Unauthorized;
