using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace dotnetCore.Controllers
{
    [Route("/[controller]")]
    [ApiController]
    public class DemoController : ControllerBase
    {
        // GET /values/
        [HttpGet]
        public ActionResult<string> Get()
        {
            return "Hello World";
        }
    }
}
