package com.atlas.dis.rest;

import com.atlas.dis.rest.processor.MonsterDropProcessor;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("monsters")
public class MonsterDropResource {
   @GET
   @Path("/drops")
   @Consumes(MediaType.APPLICATION_JSON)
   @Produces(MediaType.APPLICATION_JSON)
   public Response getDrops(@QueryParam("monsterId") Integer monsterId) {
      if (monsterId != null) {
         return MonsterDropProcessor.getByMonsterId(monsterId).build();
      }
      return MonsterDropProcessor.getAll().build();
   }
}
