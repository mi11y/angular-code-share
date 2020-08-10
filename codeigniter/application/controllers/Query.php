<?php 

class Query extends CI_Controller {

    public function __construct() {
        parent::__construct();
        $this->load->model('code_model');
        $this->load->helper('url_helper');
    }

    public function index() {
        $data['code'] = $this->code_model->get_code();
        $data['title'] = 'Code snippets';
        $this->output
            ->set_content_type('application/json')
            ->set_output(json_encode($data));
    }

}