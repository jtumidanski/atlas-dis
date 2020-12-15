package com.atlas.dis;

import java.net.URI;

import com.atlas.dis.processor.MonsterDropProcessor;
import com.atlas.shared.rest.RestServerFactory;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import database.PersistenceManager;

public class Server {
   public static void main(String[] args) {
      PersistenceManager.construct("atlas-dis");

      URI uri = UriBuilder.host(RestService.DROP_INFORMATION).uri();
      RestServerFactory.create(uri, "com.atlas.dis.rest");

      MonsterDropProcessor.init();
   }
}
